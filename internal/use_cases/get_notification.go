package usecases

import (
	"github.com/otaviouu/desafio-magalu-go/internal/core/scheduler"
)

type GetNotificationUseCase struct {
	Repo scheduler.INotificationRepository
}

func NewGetNotificationUseCase(repo scheduler.INotificationRepository) *GetNotificationUseCase {
	return &GetNotificationUseCase{
		Repo: repo,
	}
}

func (g *GetNotificationUseCase) Execute(id string) (*scheduler.Notification, error) {
	return g.Repo.GetNotificationById(id)
}
