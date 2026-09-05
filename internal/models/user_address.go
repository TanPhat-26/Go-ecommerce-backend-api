package models

import (
	"time"

	"github.com/google/uuid"
)

type UserAddress struct {
	ID             uuid.UUID `gorm:"type: uuid; default: gen_random_uuid();primaryKey"`
	UserID         uuid.UUID `gorm:"not null"`
	RecipientName  string    `gorm:"type: varchar(150); not null"`
	RecipientPhone string    `gorm:"type: varchar(30); not null"`
	AddressLine    string    `gorm:"not null"`
	City           string    `gorm:"type: varchar(100); not null"`
	District       *string   `gorm:"type: varchar(100)"`
	Ward           *string   `gorm:"type:varchar(100)"`
	PostalCode     *string   `gorm:"type:varchar(20)"`
	IsDefault      bool      `gorm:"not null; default:false"`
	CreatedAt      time.Time `gorm:"not null"`
	UpdatedAt      time.Time `gorm:"not null"`
}
