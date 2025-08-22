package routes

import (
	"expense-tracker/internal/services"

	"github.com/go-chi/chi/v5"
)

func UserRoute(userService *services.UserService) chi.Router {
	r := chi.NewRouter();
	
	r.Post("/register", userService.HandleUserRegistration);

	return r;
}