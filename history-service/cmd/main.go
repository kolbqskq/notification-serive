package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/kolbqskq/notification-service/history-service/internal/core/config"
	"github.com/kolbqskq/notification-service/history-service/internal/core/logger"
	postgres_pool "github.com/kolbqskq/notification-service/history-service/internal/core/repository/postrgers/pool"
	repository_postgres "github.com/kolbqskq/notification-service/history-service/internal/features/notification/repository/postgres"
	notification_service "github.com/kolbqskq/notification-service/history-service/internal/features/notification/service"
	transport_grpc "github.com/kolbqskq/notification-service/history-service/internal/features/notification/transport/grpc"
	pb "github.com/kolbqskq/notification-service/proto/notification/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	//Configs:
	config.Init()
	loggerConfig := config.NewLoggerConfig()
	dbConfig := config.NewDatabaseConfig()
	grpcConfig := config.NewGRPCConfig()

	//Logger:
	logger := logger.NewLogger(loggerConfig)

	//Database:
	db := postgres_pool.CreateDbPool(dbConfig.Url, logger)

	//Repositories:
	notificationRepository := repository_postgres.NewNotificationRepository(repository_postgres.NotificationRepositoryDeps{
		DbPool: db,
	})

	//Services:
	notificationService := notification_service.NewNotificationService(notification_service.NotificationServiceDeps{
		NotificationRepository: notificationRepository,
	})

	//notification server:
	notificationServer := transport_grpc.NewServer(transport_grpc.ServerDeps{
		NotificationService: notificationService,
	})

	//gRPC server:
	grpcServer := grpc.NewServer()
	pb.RegisterHistoryServiceServer(grpcServer, notificationServer)
	reflection.Register(grpcServer)

	//Listener:
	lis, err := net.Listen("tcp", grpcConfig.Addr)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to listen")
	}

	go func() {
		logger.Info().Str("addr", grpcConfig.Addr).Msg("gRPC server started")
		if err := grpcServer.Serve(lis); err != nil {
			logger.Fatal().Err(err).Msg("gRPC server error")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	logger.Info().Msg("shutting down gRPC server")
	grpcServer.GracefulStop()
}
