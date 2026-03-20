package repository_postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/kolbqskq/notification-service/history-service/internal/core/domain"
	"github.com/kolbqskq/notification-service/history-service/internal/core/errs"
)

func (r *NotificationRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.NotificationRecord, error) {
	query :=
		`
	SELECT id, user_id, event_type, message, source_service, status, created_at, sent_at, error_message
	FROM notifications
	WHERE id = @id
	`
	args := pgx.NamedArgs{
		"id": id,
	}
	res := &domain.NotificationRecord{}
	err := r.dbPool.QueryRow(ctx, query, args).Scan(&res.ID, &res.UserID, &res.Type, &res.Message, &res.SourceService, &res.Status, &res.CreatedAt, &res.SentAt, &res.ErrorMessage)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errs.ErrNotFound
		}
		return nil, err
	}
	return res, nil
}
