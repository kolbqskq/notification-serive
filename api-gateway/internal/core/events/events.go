package events

import (
	"time"

	"github.com/google/uuid"
)

const TypeNotificationCreated = "notification_created"

type NotificationCreatedPayload struct {
	UserID  uuid.UUID `json:"user_id"`
	Message string    `json:"message"`
}

type NotificationEvent struct {
	ID            uuid.UUID `json:"id"`
	Version       int       `json:"version"`
	Type          string    `json:"type"`
	Payload       any       `json:"payload"`
	SourceService string    `json:"source_service"`
	CreatedAt     time.Time `json:"created_at"`
}
