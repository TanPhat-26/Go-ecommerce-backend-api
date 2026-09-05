package dto

import (
	"testing"

	"github.com/TanPhat-26/Go-ecommerce-backend-api/internal/models"
	"github.com/google/uuid"
)

func TestToUserResponse(t *testing.T) {
	phone := "0901234567"
	userID := uuid.New()

	tests := []struct {
		name      string
		user      *models.User
		wantPhone string
		wantID    uuid.UUID
		wantEmail string
	}{
		{
			name: "user with phone",
			user: &models.User{
				ID:           userID,
				Email:        "user@example.com",
				PasswordHash: "secret-hash",
				FirstName:    "Tan",
				LastName:     "Phat",
				Phone:        &phone,
				Status:       "active",
			},
			wantPhone: phone,
			wantID:    userID,
			wantEmail: "user@example.com",
		},
		{
			name: "user without phone",
			user: &models.User{
				ID:           userID,
				Email:        "user@example.com",
				PasswordHash: "secret-hash",
				FirstName:    "Tan",
				LastName:     "Phat",
				Status:       "active",
			},
			wantPhone: "",
			wantID:    userID,
			wantEmail: "user@example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToUserResponse(tt.user)

			if got.ID != tt.wantID {
				t.Fatalf("ID = %v, want %v", got.ID, tt.wantID)
			}

			if got.Email != tt.wantEmail {
				t.Fatalf("Email = %q, want %q", got.Email, tt.wantEmail)
			}

			if got.Phone != tt.wantPhone {
				t.Fatalf("Phone = %q, want %q", got.Phone, tt.wantPhone)
			}
		})
	}
}
