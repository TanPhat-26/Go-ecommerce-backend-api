package initialize

import (
	"context"
	"fmt"

	"github.com/TanPhat-26/Go-ecommerce-backend-api/internal/models"
	"gorm.io/gorm"
)

// SeedRoles creates the default system roles
// The operation is indempotent and can be excuted multiple times safely
func SeedRoles(ctx context.Context, db *gorm.DB) error {
	defaultRoles := []string{"admin", "customer"}
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, roleName := range defaultRoles {
			var role models.Role
			err := tx.
				Where("name = ?", roleName).
				FirstOrCreate(&role, models.Role{Name: roleName}).
				Error

			if err != nil {
				return fmt.Errorf("seed role %s: %w", roleName, err)
			}
		}

		return nil
	})
}
