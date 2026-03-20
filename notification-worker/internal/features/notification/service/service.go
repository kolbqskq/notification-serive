package notification_service

import (
	"context"

	"github.com/kolbqskq/notification-service/notification-worker/internal/core/domain"
)

type NotificationRepository interface {
	Create(ctx context.Context, n *domain.NotificationRecord) error
	UpdateStatus(ctx context.Context, n *domain.NotificationRecord) error
}

type TelegramSender interface {
	Send(ctx context.Context, userID string, message string) error
}

type NotificationService struct {
	notificationRepository NotificationRepository
	telegramSender         TelegramSender
}
type NotificationServiceDeps struct {
	NotificationRepository NotificationRepository
	TelegramSender         TelegramSender
}

func NewNotificationService(deps NotificationServiceDeps) *NotificationService {
	return &NotificationService{
		notificationRepository: deps.NotificationRepository,
		telegramSender: deps.TelegramSender,
	}
}
