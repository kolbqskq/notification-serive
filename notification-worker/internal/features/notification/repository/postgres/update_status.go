package repository_postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/kolbqskq/notification-service/notification-worker/internal/core/domain"
	"github.com/kolbqskq/notification-service/notification-worker/internal/core/errs"
)

func (r *NotificationRepository) UpdateStatus(ctx context.Context, n *domain.NotificationRecord) error {
	query :=
		`
	UPDATE notifications
	SET
		status = @status,
		sent_at = @sent_at,
		error_message = @error_message
	WHERE id = @id
	`
	args := pgx.NamedArgs{
		"id":            n.ID,
		"status":        n.Status,
		"sent_at":       n.SentAt,
		"error_message": n.ErrorMessage,
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
