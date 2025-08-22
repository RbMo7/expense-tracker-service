package repositories

import (
	"context"
	"expense-tracker/internal/models"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id string) error
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	// Fixed: $4 instead of $3 for password
	query := `INSERT INTO users (id, name, email, password, is_active, created_at, updated_at) 
        VALUES ($1, $2, $3, $4, true, NOW(), NOW())`

	log.Printf("🔍 Creating user with ID: %s, Name: %s, Email: %s", user.ID, user.Name, user.Email)

	_, err := r.db.Exec(ctx, query, user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		log.Printf("❌ Database error creating user: %v", err)
		log.Printf("🔍 Query: %s", query)
		log.Printf("🔍 Parameters: ID=%s, Name=%s, Email=%s, Password len=%d",
			user.ID, user.Name, user.Email, len(user.Password))
		return err
	}

	log.Printf("✅ User created successfully: %s", user.ID)
	return nil
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
	query := `SELECT id, name, email, password, is_active, created_at, updated_at
		FROM users
		WHERE id = $1 AND is_active = true
	`

	user := &models.User{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID, &user.Name, &user.Email, &user.Password, &user.IsActive, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	query := `
        UPDATE users
        SET name = $1, email = $2, updated_at = NOW()
        WHERE id = $3 AND is_active = true`

	_, err := r.db.Exec(ctx, query, user.Name, user.Email, user.ID)
	return err
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	query := `
        UPDATE users
        SET is_active = false, updated_at = NOW()
        WHERE id = $1`

	_, err := r.db.Exec(ctx, query, id)
	return err
}
