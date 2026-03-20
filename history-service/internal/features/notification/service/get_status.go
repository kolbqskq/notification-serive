package notification_service

import (
	"context"

	"github.com/google/uuid"
	"github.com/kolbqskq/notification-service/history-service/internal/core/domain"
)

func (s *NotificationService) GetStatus(ctx context.Context, id uuid.UUID) (*domain.NotificationRecord, error) {
	return s.notificationRepository.GetByID(ctx, id)
}
