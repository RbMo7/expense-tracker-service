package models

import "time"

type User struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name" validate:"required,min=2"`
	Email     string    `json:"email" db:"email" validate:"required,email"`
	Password  string    `json:"-" db:"password"`
	IsActive  bool      `json:"is_active" db:"is_active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type UserProfile struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name" validate:"required,min=2"`
	Email     string    `json:"email" db:"email" validate:"required,email"`
	IsActive  bool      `json:"is_active" db:"is_active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type UserUpdate struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name,omitempty" db:"name" validate:"required,min=2"`
	Email     string    `json:"email,omitempty" db:"email" validate:"required,email"`
	IsActive  bool      `json:"is_active,omitempty" db:"is_active"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type UserResponse struct {
	ID        string
	Name      string
	Email     string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserCreateRequest struct {
    Name     string `json:"name" validate:"required,min=2,max=100"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=6"`
}

func (u *User) ToResponse() *UserResponse {
	return &UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		IsActive:  u.IsActive,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
