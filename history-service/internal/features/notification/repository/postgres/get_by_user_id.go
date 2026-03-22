package repository_postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/kolbqskq/notification-service/history-service/internal/core/domain"
	"github.com/kolbqskq/notification-service/history-service/internal/core/errs"
)

func (r *NotificationRepository) GetByUserID(ctx context.Context, userID uuid.UUID, limit, offset int32) ([]*domain.NotificationRecord, int32, error) {
	query :=
		`
	SELECT id, user_id, event_type, message, source_service, status, created_at, sent_at, error_message
	FROM notifications
	WHERE user_id = @user_id
	ORDER BY created_at DESC
	LIMIT @limit OFFSET @offset
	`

	args := pgx.NamedArgs{
		"user_id": userID,
		"limit":   limit,
		"offset":  offset,
	}

	rows, err := r.dbPool.Query(ctx, query, args)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var res []*domain.NotificationRecord

	for rows.Next() {
		r := &domain.NotificationRecord{}
		err := rows.Scan(&r.ID, &r.UserID, &r.Type, &r.Message, &r.SourceService, &r.Status, &r.CreatedAt, &r.SentAt, &r.ErrorMessage)
		if err != nil {
			return nil, 0, err
		}
		res = append(res, r)
	}

	if len(res) == 0 {
		return nil, 0, errs.ErrNotFound
	}

	query =
		`
	SELECT COUNT (*) FROM notifications WHERE user_id = @user_id
	`

	var total int32

	err = r.dbPool.QueryRow(ctx, query, args).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return res, total, nil
}
