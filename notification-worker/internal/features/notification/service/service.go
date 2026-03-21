package notification_service

import (
	"context"

	"github.com/kolbqskq/notification-service/notification-worker/internal/core/domain"
	"github.com/rs/zerolog"
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
	logger                 zerolog.Logger
}
type NotificationServiceDeps struct {
	NotificationRepository NotificationRepository
	TelegramSender         TelegramSender
	Logger                 zerolog.Logger
}

func NewNotificationService(deps NotificationServiceDeps) *NotificationService {
	return &NotificationService{
		notificationRepository: deps.NotificationRepository,
		telegramSender:         deps.TelegramSender,
		logger:                 deps.Logger,
	}
}
