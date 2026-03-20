package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/kolbqskq/notification-service/notification-worker/internal/core/errs"
)

type NotificationStatus string

const (
	StatusPending NotificationStatus = "pending"
	StatusSent    NotificationStatus = "sent"
	StatusFailed  NotificationStatus = "failed"
)

type NotificationRecord struct {
	ID            uuid.UUID
	UserID        uuid.UUID
	Type          string
	Message       string
	SourceService string
	Status        NotificationStatus
	CreatedAt     time.Time
	SentAt        *time.Time
	ErrorMessage  *string
}

func NewNotificationRecord(userID uuid.UUID, t string, message string, source string) (*NotificationRecord, error) {
	if message == "" {
		return nil, errs.ErrEmptyMessage
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return &NotificationRecord{
		ID:            id,
		UserID:        userID,
		Type:          t,
		Message:       message,
		SourceService: source,
		Status:        StatusPending,
		CreatedAt:     time.Now(),
	}, nil
}

func (n *NotificationRecord) MarkSend() {
	now := time.Now()
	n.Status = StatusSent
	n.SentAt = &now
}

func (n *NotificationRecord) MarkFailed(err string) {
	n.Status = StatusFailed
	n.ErrorMessage = &err
}
