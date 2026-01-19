package scheduler

import (
	"time"

	"github.com/otaviouu/desafio-magalu-go/internal/use_cases/dtos"
)

type INotificationRepository interface {
	GetNotificationById(id string) (*Notification, error)
	CreateNotification(notificationInput *dtos.NotificationInput) (*Notification, error)
}

type Notification struct {
	Id      string    `db:"id"`
	Message string    `db:"message"`
	SendAt  time.Time `db:"send_at"`
	Status  string    `db:"status"`
}

func (n *Notification) IsValid() bool {
	if n.Id == "" || n.Message == "" || n.SendAt.IsZero() {
		return false
	}

	if !is_status_valid(n.Status) {
		return false

	}

	return true
}

func is_status_valid(status string) bool {
	if status == "pending" || status == "sent" || status == "failed" {
		return true
	}
	return false
}
