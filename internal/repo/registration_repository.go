package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/TanPhat-26/Go-ecommerce-backend-api/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

// ErrEmailAlreadyExists indicates that the user email index rejected a duplicate value
var ErrEmailAlreadyExists = errors.New("email already exists")

// RegistrationRepository persists the atomic user-registration workflow
type RegistrationRepository interface {
	CreateUserWithRole(
		ctx context.Context,
		user *models.User,
		roleName string,
	) error
}

type registrationRepository struct {
	db *gorm.DB
}

func NewRegistrationRepository(
	db *gorm.DB,
) RegistrationRepository {
	return &registrationRepository{
		db: db,
	}
}

// CreateUserWithRole creates a user and assigns a role in one transaction
func (r *registrationRepository) CreateUserWithRole(
	ctx context.Context,
	user *models.User,
	roleName string,
) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var role models.Role

		if err := tx.
			Where("name = ?", roleName).
			First(&role).Error; err != nil {
			return fmt.Errorf("find role %q: %w", roleName, err)
		}

		if user.ID == uuid.Nil {
			user.ID = uuid.New()
		}

		if err := tx.Create(user).Error; err != nil {
			if isUniqueViolation(err, "ux_users_email_lower") {
				return ErrEmailAlreadyExists
			}

			return fmt.Errorf("create user: %w", err)
		}

		userRole := &models.UserRole{
			UserID: user.ID,
			RoleID: role.ID,
		}

		if err := tx.Create(userRole).Error; err != nil {
			return fmt.Errorf("assign role %q: %w", roleName, err)
		}

		return nil
	})
}

func isUniqueViolation(err error, constraintName string) bool {
	var postgresError *pgconn.PgError

	return errors.As(err, &postgresError) &&
		postgresError.Code == "23505" &&
		postgresError.ConstraintName == constraintName
}
