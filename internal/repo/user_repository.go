package repo

import (
	"context"

	"github.com/TanPhat-26/Go-ecommerce-backend-api/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserRepository defines database operation for users
// It abstracts persistence details from the service layer
type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
}

// userRepository is the GORM-based implementation of UserRepository
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a user repository with the provided database
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(
	ctx context.Context,
	user *models.User,
) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepository) FindByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.User, error) {
	var user models.User

	err := r.db.
		WithContext(ctx).
		Where("id = ?", id).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// FindByEmail retrieves a user by email without case sensitivity
// The query uses the database index created for normalized email lookup
func (r *userRepository) FindByEmail(
	ctx context.Context,
	email string,
) (*models.User, error) {
	var user models.User

	err := r.db.
		WithContext(ctx).
		Where("LOWER(email) = LOWER(?)", email).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Update(
	ctx context.Context,
	user *models.User,
) error {
	return r.db.WithContext(ctx).Save(user).Error
}
