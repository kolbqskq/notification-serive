package domain

import (
	"github.com/google/uuid"
	"github.com/kolbqskq/notification-service/api-gateway/internal/core/errs"
)

type Notification struct {
	UserID  uuid.UUID
	Message string
}

func NewNotification(userID uuid.UUID, message string) (*Notification, error) {
	if message == "" {
		return nil, errs.ErrEmptyMessage
	}
	return &Notification{
		UserID:  userID,
		Message: message,
	}, nil
}
