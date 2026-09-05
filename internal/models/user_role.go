package models

import (
	"time"

	"github.com/google/uuid"
)

type UserRole struct {
	UserID    uuid.UUID `gorm:"type:uuid;not null;primaryKey"`
	RoleID    uuid.UUID `gorm:"type:uuid;not null;primaryKey"`
	CreatedAt time.Time `gorm:"not null"`
}
