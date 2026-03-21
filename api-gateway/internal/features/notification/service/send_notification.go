package notification_service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/kolbqskq/notification-service/api-gateway/internal/core/domain"
	"github.com/kolbqskq/notification-service/api-gateway/internal/core/events"
)

func (s *NotificationService) SendNotification(ctx context.Context, n *domain.Notification) (uuid.UUID, error) {
	payload := events.NotificationCreatedPayload{
		UserID:  n.UserID,
		Message: n.Message,
	}

	id, err := uuid.NewV7()
	if err != nil {
		return uuid.Nil, err
	}

	event := events.NotificationEvent{
		ID:            id,
		Version:       1,
		Type:          events.TypeNotificationCreated,
		Payload:       payload,
		SourceService: s.serviceName,
		CreatedAt:     time.Now(),
	}
	if err := s.producer.Publish(ctx, n.UserID.String(), event); err != nil {
		return uuid.Nil, err
	}
	s.logger.Info().Str("user_id", n.UserID.String()).Str("notification_id", event.ID.String()).Msg("notification sent")
	return id, nil
}
