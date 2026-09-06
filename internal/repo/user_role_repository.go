package repo

import (
	"context"

	"github.com/TanPhat-26/Go-ecommerce-backend-api/internal/models"
	"gorm.io/gorm"
)

// UserRoleRepository defines database operations for user-role assignments
type UserRoleRepository interface {
	Create(ctx context.Context, userRole *models.UserRole) error
}

type userRoleRepository struct {
	db *gorm.DB
}

func NewUserRoleRepository(db *gorm.DB) UserRoleRepository {
	return &userRoleRepository{
		db: db,
	}
}

func (r *userRoleRepository) Create(
	ctx context.Context,
	userRole *models.UserRole,
) error {
	return r.db.WithContext(ctx).Create(userRole).Error
}
