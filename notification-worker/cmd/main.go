package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/kolbqskq/notification-service/notification-worker/internal/core/config"
	"github.com/kolbqskq/notification-service/notification-worker/internal/core/logger"
	postgres_pool "github.com/kolbqskq/notification-service/notification-worker/internal/core/repository/postrgers/pool"
	"github.com/kolbqskq/notification-service/notification-worker/internal/core/transport/kafka"
	repository_postgres "github.com/kolbqskq/notification-service/notification-worker/internal/features/notification/repository/postgres"
	notification_service "github.com/kolbqskq/notification-service/notification-worker/internal/features/notification/service"
	transport_kafka "github.com/kolbqskq/notification-service/notification-worker/internal/features/notification/transport/kafka"
	transport_telegram "github.com/kolbqskq/notification-service/notification-worker/internal/features/notification/transport/telegram"
)

func main() {
	//Configs:
	config.Init()
	loggerConfig := config.NewLoggerConfig()
	kafkaConfig := config.NewKafkaConfig()
	dbConfig := config.NewDatabaseConfig()
	telegramConfig := config.NewTelegramConfig()

	//Logger:
	logger := logger.NewLogger(loggerConfig)

	//Database:
	db := postgres_pool.CreateDbPool(dbConfig.Url, logger)

	//Telegram:
	telegramClient, err := transport_telegram.NewTelegramClient(transport_telegram.TelegramClientDeps{
		Token:  telegramConfig.Token,
		ChatID: telegramConfig.ChatID,
	})
	if err != nil {
		logger.Fatal().Err(err).Msg("failed create telegram client")
	}

	//Repositories:
	notificationRepository := repository_postgres.NewNotificationRepository(repository_postgres.NotificationRepositoryDeps{
		DbPool: db,
	})

	//Services:
	notificationService := notification_service.NewNotificationService(notification_service.NotificationServiceDeps{
		NotificationRepository: notificationRepository,
		TelegramSender:         telegramClient,
		Logger:                 *logger,
	})

	//Handler:
	handler := transport_kafka.NewNotificationHandler(transport_kafka.NotificationHandleDeps{
		NotificationService: notificationService,
	})

	//Consumer:
	consumer := kafka.NewKafkaConsumer(kafka.KafkaConsumerDeps{
		Brokers: kafkaConfig.Brokers,
		Topic:   kafkaConfig.Topic,
		GroupID: kafkaConfig.GroupID,
		Handler: handler.Handle,
		Logger:  *logger,
	})

	logger.Info().Strs("brokers", kafkaConfig.Brokers).Str("topic", kafkaConfig.Topic).Msg("consumer created")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		defer func() {
			if err := consumer.Close(); err != nil {
				logger.Error().Err(err).Msg("error closing consumer")
			}
		}()
		if err := consumer.Run(ctx); err != nil {
			logger.Fatal().Err(err).Msg("consumer stopped")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	cancel()
}
