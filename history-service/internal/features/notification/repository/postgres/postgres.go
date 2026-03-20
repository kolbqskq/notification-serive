package repository_postgres

import "github.com/jackc/pgx/v5/pgxpool"

type NotificationRepository struct {
	dbPool *pgxpool.Pool
}

type NotificationRepositoryDeps struct {
	DbPool *pgxpool.Pool
}

func NewNotificationRepository(deps NotificationRepositoryDeps) *NotificationRepository {
	return &NotificationRepository{
		dbPool: deps.DbPool,
	}
}
