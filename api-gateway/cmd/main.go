package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kolbqskq/notification-service/api-gateway/internal/core/config"
	"github.com/kolbqskq/notification-service/api-gateway/internal/core/logger"
	transport_grpc "github.com/kolbqskq/notification-service/api-gateway/internal/core/transport/grpc"
	transport_kafka "github.com/kolbqskq/notification-service/api-gateway/internal/core/transport/kafka"
	"github.com/kolbqskq/notification-service/api-gateway/internal/core/transport/middleware"
	notification_service "github.com/kolbqskq/notification-service/api-gateway/internal/features/notification/service"
	notification_http "github.com/kolbqskq/notification-service/api-gateway/internal/features/notification/transport/http"
	"github.com/redis/go-redis/v9"
)

func main() {
	//Configs:
	config.Init()
	loggerConfig := config.NewLoggerConfig()
	redisConfig := config.NewAddrConfig()
	kafkaConfig := config.NewKafkaConfig()
	serviceConfig := config.NewServiceConfig()
	appConfig := config.NewAppConfig()
	historyServiceConfig := config.NewHistoryServiceConfig()

	//Logger:
	logger := logger.NewLogger(loggerConfig)
	//Redis:
	rdb := redis.NewClient(&redis.Options{
		Addr: redisConfig.RedisAddr,
	})
	defer rdb.Close()

	//Producers:
	kafkaProducer := transport_kafka.NewKafkaProducer(kafkaConfig.Brokers, kafkaConfig.Topic)
	defer kafkaProducer.Close()

	app := gin.New()
	app.Use(gin.Recovery())
	app.SetTrustedProxies(nil)
	app.Use(middleware.ErrorMiddleware(*logger))
	app.Use(middleware.RateLimit(rdb))

	//Services:
	notificationService := notification_service.NewNotificationService(kafkaProducer, serviceConfig.Name)

	//HistoryClient:
	historyClient, conn, err := transport_grpc.NewHistoryClient(historyServiceConfig.Addr)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to connect history service")
	}
	defer conn.Close()

	//Handlers:

	notification_http.NewNotificationHandler(notification_http.NotificationHandlerDeps{
		Router:              app,
		NotificationService: notificationService,
		HistoryClient:       historyClient,
	})

	server := &http.Server{
		Addr:    appConfig.Addr,
		Handler: app,
	}

	go func() {
		logger.Info().Str("addr", server.Addr).Msg("server started")

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal().Err(err).Msg("server error")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal().Err(err).Msg("server shutdown error")
	}

}
