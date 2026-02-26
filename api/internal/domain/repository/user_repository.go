package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/next-go-template/api/internal/domain/entity"
)

// UserRepository defines persistence operations for users.
type UserRepository interface {
	Create(ctx context.Context, u *entity.User) error
	GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	Update(ctx context.Context, u *entity.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}
