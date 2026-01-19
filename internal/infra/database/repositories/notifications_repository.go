package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/otaviouu/desafio-magalu-go/internal/core/scheduler"
)

type NotificationRepository struct {
	db *sqlx.DB
}

func NewNotificationRepository(db *sqlx.DB) (*NotificationRepository, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	return &NotificationRepository{
		db: db,
	}, nil
}

func (nr *NotificationRepository) GetNotificationById(id string) (*scheduler.Notification, error) {
	var notification scheduler.Notification
	err := nr.db.Get(&notification, "SELECT id, message, send_at, status FROM notifications WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return &notification, nil
}
