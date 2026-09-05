package repo

import (
	"context"
	"time"

	"github.com/TanPhat-26/Go-ecommerce-backend-api/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RefreshTokenRepository manages refresh token persistence and revocation.
type RefreshTokenRepository interface {
	Create(ctx context.Context, token *models.RefreshToken) error
	FindByHash(ctx context.Context, tokenHash string) (*models.RefreshToken, error)
	Revoke(ctx context.Context, id uuid.UUID) error
	RevokeAllByUserID(ctx context.Context, userID uuid.UUID) error
}

// refreshTokenRepository is the GORM-based implementation of RefreshTokenRepository.
type refreshTokenRepository struct {
	db *gorm.DB
}

// NewRefreshTokenRepository creates a refresh token repository.
func NewRefreshTokenRepository(db *gorm.DB) RefreshTokenRepository {
	return &refreshTokenRepository{
		db: db,
	}
}

func (r *refreshTokenRepository) Create(
	ctx context.Context,
	token *models.RefreshToken,
) error {
	return r.db.WithContext(ctx).Create(token).Error
}

func (r *refreshTokenRepository) FindByHash(
	ctx context.Context,
	tokenHash string,
) (*models.RefreshToken, error) {
	var token models.RefreshToken

	err := r.db.
		WithContext(ctx).
		Where("token_hash = ?", tokenHash).
		First(&token).
		Error

	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (r *refreshTokenRepository) Revoke(
	ctx context.Context,
	id uuid.UUID,
) error {
	return r.db.
		WithContext(ctx).
		Model(&models.RefreshToken{}).
		Where("id = ? AND revoked_at IS NULL", id).
		Update("revoked_at", time.Now().UTC()).
		Error
}

func (r *refreshTokenRepository) RevokeAllByUserID(
	ctx context.Context,
	userID uuid.UUID,
) error {
	return r.db.
		WithContext(ctx).
		Model(&models.RefreshToken{}).
		Where("user_id = ? AND revoked_at IS NULL", userID).
		Update("revoked_at", time.Now().UTC()).
		Error
}
