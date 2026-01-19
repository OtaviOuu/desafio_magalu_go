package dtos

import (
	"time"
)

type NotificationInput struct {
	Message string
	SendAt  time.Time
	Status  string
}

func (nin *NotificationInput) IsValid() bool {
	if nin.Message == "" || nin.SendAt.IsZero() || nin.Status == "" {
		return false
	}
	return true
}

type NotificationOutput struct {
	ID      string
	Message string
	SendAt  time.Time
	Status  string
}
