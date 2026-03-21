package transport_grpc

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/kolbqskq/notification-service/history-service/internal/core/domain"
	"github.com/kolbqskq/notification-service/history-service/internal/core/errs"
	pb "github.com/kolbqskq/notification-service/proto/notification/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NotificationService interface {
	GetHistory(ctx context.Context, userID uuid.UUID, limit, offset int32) ([]*domain.NotificationRecord, int32, error)
	GetStatus(ctx context.Context, id uuid.UUID) (*domain.NotificationRecord, error)
}

type Server struct {
	pb.UnimplementedHistoryServiceServer
	notificationService NotificationService
}

type ServerDeps struct {
	NotificationService NotificationService
}

func NewServer(deps ServerDeps) *Server {
	return &Server{
		notificationService: deps.NotificationService,
	}
}

func toGRPCError(err error) error {
	switch err {
	case errs.ErrNotFound:
		return status.Error(codes.NotFound, err.Error())
	case errs.ErrInvalidID:
		return status.Error(codes.InvalidArgument, err.Error())
	default:
		fmt.Printf("unknown error: %v\n", err)
		return status.Error(codes.Internal, "internal error")
	}
}
