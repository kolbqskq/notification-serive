package transport_kafka

import (
	"context"
	"encoding/json"

	"github.com/kolbqskq/notification-service/notification-worker/internal/core/domain"
	"github.com/kolbqskq/notification-service/notification-worker/internal/core/events"
)

type NotificationService interface {
	HandleNotification(ctx context.Context, record *domain.NotificationRecord) error
}

type NotificationHandler struct {
	notificationService NotificationService
}

type NotificationHandleDeps struct {
	NotificationService NotificationService
}

func NewNotificationHandler(deps NotificationHandleDeps) *NotificationHandler {
	return &NotificationHandler{
		notificationService: deps.NotificationService,
	}
}

func (h *NotificationHandler) Handle(ctx context.Context, event events.NotificationEvent) error {
	switch event.Type {
	case events.TypeNotificationCreated:
		var payload events.NotificationCreatedPayload

		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			return err
		}

		record, err := domain.NewNotificationRecord(event.ID, payload.UserID, event.Type, payload.Message, event.SourceService, event.CreatedAt)
		if err != nil {
			return err
		}

		return h.notificationService.HandleNotification(ctx, record)

	default:
		return nil
	}
}
