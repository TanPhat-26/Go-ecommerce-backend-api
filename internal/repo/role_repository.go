package repo

import (
	"context"

	"github.com/TanPhat-26/Go-ecommerce-backend-api/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RoleRepository defines database operations for roles.
type RoleRepository interface {
	Create(ctx context.Context, role *models.Role) error
	FindByID(ctx context.Context, id uuid.UUID) (*models.Role, error)
	FindByName(ctx context.Context, name string) (*models.Role, error)
	List(ctx context.Context) ([]models.Role, error)
}

// roleRepository is the GORM-based implementation of RoleRepository.
type roleRepository struct {
	db *gorm.DB
}

// NewRoleRepository creates a role repository with the provided database connection.
func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		db: db,
	}
}

func (r *roleRepository) Create(
	ctx context.Context,
	role *models.Role,
) error {
	return r.db.WithContext(ctx).Create(role).Error
}

func (r *roleRepository) FindByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.Role, error) {
	var role models.Role

	err := r.db.
		WithContext(ctx).
		Where("id = ?", id).
		First(&role).
		Error

	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *roleRepository) FindByName(
	ctx context.Context,
	name string,
) (*models.Role, error) {
	var role models.Role

	err := r.db.
		WithContext(ctx).
		Where("LOWER(name) = LOWER(?)", name).
		First(&role).
		Error

	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *roleRepository) List(
	ctx context.Context,
) ([]models.Role, error) {
	var roles []models.Role

	err := r.db.
		WithContext(ctx).
		Order("name ASC").
		Find(&roles).
		Error

	if err != nil {
		return nil, err
	}

	return roles, nil
}
