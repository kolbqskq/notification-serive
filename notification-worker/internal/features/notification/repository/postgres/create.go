package repository_postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/kolbqskq/notification-service/notification-worker/internal/core/domain"
	"github.com/kolbqskq/notification-service/notification-worker/internal/core/errs"
)

func (r *NotificationRepository) Create(ctx context.Context, n *domain.NotificationRecord) error {
	query :=
		`
	INSERT INTO notifications (id, user_id, event_type, message, source_service, status, created_at)
	VALUES(@id, @user_id, @event_type, @message, @source_service, @status, @created_at)
	`
	args := pgx.NamedArgs{
		"id":             n.ID,
		"user_id":        n.UserID,
		"event_type":     n.Type,
		"message":        n.Message,
		"source_service": n.SourceService,
		"status":         n.Status,
		"created_at":     n.CreatedAt,
	}

	row, err := r.dbPool.Exec(ctx, query, args)
	if err != nil {
		return err
	}
	
	if row.RowsAffected() == 0 {
		return errs.ErrNotFound
	}

	return nil
}
