package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/otaviouu/desafio-magalu-go/internal/core/scheduler"
	usecases "github.com/otaviouu/desafio-magalu-go/internal/use_cases"
	"github.com/otaviouu/desafio-magalu-go/internal/use_cases/dtos"
)

type NotificationsHandler struct {
	Repo scheduler.INotificationRepository
}

func NewNotificationsHandler(repo scheduler.INotificationRepository) *NotificationsHandler {
	return &NotificationsHandler{
		Repo: repo,
	}
}

func (h *NotificationsHandler) CreateNotification(w http.ResponseWriter, r *http.Request) {
	var notificationInput *dtos.NotificationInput
	if err := json.NewDecoder(r.Body).Decode(&notificationInput); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createNotificationUseCase := usecases.NewCreateNotificationUseCase(h.Repo)

	notification, err := createNotificationUseCase.Execute(notificationInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	notificationOutput := &dtos.NotificationOutput{
		ID:      notification.ID,
		Message: notification.Message,
		SendAt:  notification.SendAt,
		Status:  notification.Status,
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(notificationOutput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *NotificationsHandler) GetNotificationById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "wrong id", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	getNotificationUseCase := usecases.NewGetNotificationUseCase(h.Repo)

	notification, err := getNotificationUseCase.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	status := string(notification.Status)

	notificationOutput := &dtos.NotificationOutput{
		ID:      notification.Id,
		Message: notification.Message,
		SendAt:  notification.SendAt,
		Status:  status,
	}

	err = json.NewEncoder(w).Encode(notificationOutput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
