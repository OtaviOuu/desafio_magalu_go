package usecases

import (
	"github.com/otaviouu/desafio-magalu-go/internal/core/scheduler"
	"github.com/otaviouu/desafio-magalu-go/internal/use_cases/dtos"
)

type CreateNotificationUseCase struct {
	Repo scheduler.INotificationRepository
}

func NewCreateNotificationUseCase(repo scheduler.INotificationRepository) *CreateNotificationUseCase {
	return &CreateNotificationUseCase{
		Repo: repo,
	}
}

func (c *CreateNotificationUseCase) Execute(notificationRequest *dtos.NotificationInput) (*dtos.NotificationOutput, error) {
	notification, err := c.Repo.CreateNotification(notificationRequest)
	if err != nil {
		return nil, err
	}

	return &dtos.NotificationOutput{
		ID:      notification.Id,
		Message: notification.Message,
		SendAt:  notification.SendAt,
		Status:  notification.Status,
	}, nil
}
