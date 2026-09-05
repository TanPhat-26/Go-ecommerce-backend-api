package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email        string         `gorm:"type:varchar(255);not null"`
	PasswordHash string         `gorm:"column:password_hash;not null"`
	FirstName    string         `gorm:"type:varchar(100);not null"`
	LastName     string         `gorm:"type:varchar(100);not null"`
	Phone        *string        `gorm:"type:varchar(30)"`
	Status       string         `gorm:"type:varchar(30);not null;default:active"`
	CreatedAt    time.Time      `gorm:"not null"`
	UpdatedAt    time.Time      `gorm:"not null"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
