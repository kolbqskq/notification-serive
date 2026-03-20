package notification_service

import "context"

type Producer interface {
	Publish(ctx context.Context, key string, v any) error
}

type NotificationService struct {
	producer    Producer
	serviceName string
}

func NewNotificationService(producer Producer, serviceName string) *NotificationService {
	return &NotificationService{
		producer:    producer,
		serviceName: serviceName,
	}
}
