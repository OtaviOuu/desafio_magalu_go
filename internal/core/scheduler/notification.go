package scheduler

import (
	"time"
)

type INotificationRepository interface {
	GetNotificationById(id string) (*Notification, error)
}

type SendStatus string

const (
	Pending SendStatus = "pending"
	Sent    SendStatus = "sent"
	Failed  SendStatus = "failed"
)

type Notification struct {
	Id      string     `db:"id"`
	Message string     `db:"message"`
	SendAt  time.Time  `db:"send_at"`
	Status  SendStatus `db:"status"`
}

func (n *Notification) IsValid() bool {
	if n.Id == "" || n.Message == "" || n.SendAt.IsZero() || n.Status == "" {
		return false
	}
	return true
}
