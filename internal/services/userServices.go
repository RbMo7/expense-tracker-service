package services

import (
	"context"
	"encoding/json"
	"expense-tracker/internal/models"
	"expense-tracker/internal/repositories"
	"expense-tracker/internal/workers"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) HandleUserRegistration(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

    var req models.UserCreateRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        workers.SendJSONError(w, "Invalid JSON format", "json_parse_error", http.StatusBadRequest)
        return
    }

    // Validate required fields
    if req.Name == "" || req.Email == "" || req.Password == "" {
        workers.SendJSONError(w, "Name, email, and password are required", "validation_error", http.StatusBadRequest)
        return
    }

    // Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        workers.SendJSONError(w, "Failed to process password", "server_error", http.StatusInternalServerError)
        return
    }

    // Create user model
    user := &models.User{
        ID:       uuid.New().String(),
        Name:     req.Name,
        Email:    req.Email,
        Password: string(hashedPassword),
        IsActive: true,
    }

    // Save to database
    if err := s.userRepo.Create(ctx, user); err != nil {
        workers.SendJSONError(w, "Failed to create user", "database_error", http.StatusInternalServerError)
        return
    }

    // Return success response
    workers.SendJSONSuccess(w, "User created successfully", user.ToResponse(), http.StatusCreated)
}

func HandleUserLogin(w http.ResponseWriter, r *http.Request) {
	// Handle user login logic here
	w.Write([]byte("User logged in successfully!"))
}

func HandleUserLogout(w http.ResponseWriter, r *http.Request) {
	// Handle user logout logic here
}

func HandleUserProfile(w http.ResponseWriter, r *http.Request) {
	// Handle user profile logic here
}

func HandleUserUpdate(w http.ResponseWriter, r *http.Request) {
	// Handle user update logic here
}

func HandleUserDeletion(w http.ResponseWriter, r *http.Request) {
	// Handle user deletion logic here
}

