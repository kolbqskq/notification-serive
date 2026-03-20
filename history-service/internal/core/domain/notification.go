package domain

import (
	"time"

	"github.com/google/uuid"
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
