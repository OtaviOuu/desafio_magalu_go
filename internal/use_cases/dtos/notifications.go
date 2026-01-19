package dtos

import (
	"time"

	"github.com/otaviouu/desafio-magalu-go/internal/core/scheduler"
)

type NotificationInput struct {
	Message string               `json:"message"`
	SendAt  time.Time            `json:"send_at"`
	Status  scheduler.SendStatus `json:"status"`
}

type NotificationOutput struct {
	ID      string               `json:"id"`
	Message string               `json:"message"`
	SendAt  time.Time            `json:"send_at"`
	Status  scheduler.SendStatus `json:"status"`
}
