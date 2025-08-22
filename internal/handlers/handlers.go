package handlers

import (
	"expense-tracker/internal/database"
	"expense-tracker/internal/repositories"
	"expense-tracker/internal/routes"
	"expense-tracker/internal/services"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r *chi.Mux) {
	// Test connection route
	r.Get("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong!"))
	})

	userRepo := repositories.NewUserRepository(database.DB)
	userService := services.NewUserService(userRepo)

	r.Mount("/api/user", routes.UserRoute(userService))
}
