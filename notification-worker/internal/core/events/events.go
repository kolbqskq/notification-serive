package events

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

const TypeNotificationCreated = "notification_created"

type NotificationCreatedPayload struct {
	UserID  uuid.UUID `json:"user_id"`
	Message string    `json:"message"`
}

type NotificationEvent struct {
	ID            uuid.UUID       `json:"id"`
	Version       int             `json:"version"`
	Type          string          `json:"type"`
	Payload       json.RawMessage `json:"payload"`
	SourceService string          `json:"source_service"`
	CreatedAt     time.Time       `json:"created_at"`
}
