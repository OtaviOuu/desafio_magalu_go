package repositories

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/otaviouu/desafio-magalu-go/internal/core/scheduler"
	"github.com/otaviouu/desafio-magalu-go/internal/use_cases/dtos"
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

func (nr *NotificationRepository) CreateNotification(notificationInput *dtos.NotificationInput) (*scheduler.Notification, error) {
	uuid := uuid.New().String()

	query := `
		INSERT INTO notifications (id, message, send_at, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id, message, send_at, status
	`

	var n scheduler.Notification
	err := nr.db.QueryRowx(query, uuid, notificationInput.Message, notificationInput.SendAt, notificationInput.Status).StructScan(&n)
	if err != nil {
		return nil, err
	}

	return &n, nil
}
