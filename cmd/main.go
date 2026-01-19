package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/otaviouu/desafio-magalu-go/internal/handlers"
	"github.com/otaviouu/desafio-magalu-go/internal/infra/database/repositories"
)

func main() {

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db, err := sqlx.Connect(
		"postgres",
		"host=localhost port=5434 user=postgres password=postgres dbname=postgres sslmode=disable",
	)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	notificationsRepo, err := repositories.NewNotificationRepository(db)
	if err != nil {
		panic(err)
	}

	notificationHandler := handlers.NewNotificationsHandler(notificationsRepo)

	r.Route("/notifications", func(r chi.Router) {
		r.Post("/", notificationHandler.CreateNotification)
		r.Get("/{id}", notificationHandler.GetNotificationById)
	})

	log.Println("server...")

	err = http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal(err)
	}

}
