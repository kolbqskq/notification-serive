package notification_service

import (
	"context"

	"github.com/google/uuid"
	"github.com/kolbqskq/notification-service/history-service/internal/core/domain"
)

type NotificationRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*domain.NotificationRecord, error)
	GetByUserID(ctx context.Context, userID uuid.UUID, limit, offset int32) ([]*domain.NotificationRecord, int32, error)
}

type NotificationService struct {
	notificationRepository NotificationRepository
}

type NotificationServiceDeps struct {
	NotificationRepository NotificationRepository
}

func NewNotificationService(deps NotificationServiceDeps) *NotificationService {
	return &NotificationService{
		notificationRepository: deps.NotificationRepository,
	}
}
