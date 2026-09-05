package repo

import (
	"context"
	"time"

	"github.com/TanPhat-26/Go-ecommerce-backend-api/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserAddressRepository defines database operations for user addresses
type UserAddressRepository interface {
	Create(ctx context.Context, address *models.UserAddress) error
	FindByIDForUser(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*models.UserAddress, error)
	ListByUserID(ctx context.Context, userID uuid.UUID) ([]models.UserAddress, error)
	FindDefaultByUserID(ctx context.Context, userID uuid.UUID) (*models.UserAddress, error)
	UpdateByIDForUser(ctx context.Context, id uuid.UUID, userID uuid.UUID, address *models.UserAddress) error
	DeleteByIDForUser(ctx context.Context, id uuid.UUID, userID uuid.UUID) error
}

// userAddressRepository is the GORM-based implementation of UserAddressRepository.
type userAddressRepository struct {
	db *gorm.DB
}

// NewUserAddressRepository creates a user address repository.
func NewUserAddressRepository(db *gorm.DB) UserAddressRepository {
	return &userAddressRepository{
		db: db,
	}
}

func (r *userAddressRepository) Create(
	ctx context.Context,
	address *models.UserAddress,
) error {
	return r.db.WithContext(ctx).Create(address).Error
}

func (r *userAddressRepository) FindByIDForUser(
	ctx context.Context,
	id uuid.UUID,
	userID uuid.UUID,
) (*models.UserAddress, error) {
	var address models.UserAddress

	err := r.db.
		WithContext(ctx).
		Where("id = ? AND user_id = ?", id, userID).
		First(&address).
		Error

	if err != nil {
		return nil, err
	}

	return &address, nil
}

func (r *userAddressRepository) ListByUserID(
	ctx context.Context,
	userID uuid.UUID,
) ([]models.UserAddress, error) {
	var addresses []models.UserAddress

	err := r.db.
		WithContext(ctx).
		Where("user_id = ?", userID).
		Order("is_default DESC, created_at DESC").
		Find(&addresses).
		Error

	if err != nil {
		return nil, err
	}

	return addresses, nil
}

func (r *userAddressRepository) FindDefaultByUserID(
	ctx context.Context,
	userID uuid.UUID,
) (*models.UserAddress, error) {
	var address models.UserAddress

	err := r.db.
		WithContext(ctx).
		Where("user_id = ? AND is_default = ?", userID, true).
		First(&address).
		Error

	if err != nil {
		return nil, err
	}

	return &address, nil
}

func (r *userAddressRepository) UpdateByIDForUser(
	ctx context.Context,
	id uuid.UUID,
	userID uuid.UUID,
	address *models.UserAddress,
) error {
	result := r.db.
		WithContext(ctx).
		Model(&models.UserAddress{}).
		Where("id = ? AND user_id = ?", id, userID).
		Updates(map[string]interface{}{
			"recipient_name":  address.RecipientName,
			"recipient_phone": address.RecipientPhone,
			"address_line":    address.AddressLine,
			"city":            address.City,
			"district":        address.District,
			"ward":            address.Ward,
			"postal_code":     address.PostalCode,
			"is_default":      address.IsDefault,
			"updated_at":      time.Now().UTC(),
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *userAddressRepository) DeleteByIDForUser(
	ctx context.Context,
	id uuid.UUID,
	userID uuid.UUID,
) error {
	result := r.db.
		WithContext(ctx).
		Where("id = ? AND user_id = ?", id, userID).
		Delete(&models.UserAddress{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}