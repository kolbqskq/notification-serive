package notification_service

import (
	"context"

	"github.com/google/uuid"
	"github.com/kolbqskq/notification-service/history-service/internal/core/domain"
)

const (
	defaultLimit = 20
	maxLimit     = 100
)

func (s *NotificationService) GetHistory(ctx context.Context, userID uuid.UUID, limit, offset int32) ([]*domain.NotificationRecord, int32, error) {
	if limit == 0 {
		limit = defaultLimit
	}
	if limit > maxLimit {
		limit = maxLimit
	}
	if offset < 0 {
		offset = 0
	}

	return s.notificationRepository.GetByUserID(ctx, userID, limit, offset)
}
