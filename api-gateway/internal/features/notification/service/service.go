package notification_service

import (
	"context"

	"github.com/rs/zerolog"
)

type Producer interface {
	Publish(ctx context.Context, key string, v any) error
}

type NotificationService struct {
	producer    Producer
	serviceName string
	logger      zerolog.Logger
}

type NotificationServiceDeps struct {
	Producer    Producer
	ServiceName string
	Logger      zerolog.Logger
}

func NewNotificationService(deps NotificationServiceDeps) *NotificationService {
	return &NotificationService{
		producer:    deps.Producer,
		serviceName: deps.ServiceName,
		logger:      deps.Logger,
	}
}
