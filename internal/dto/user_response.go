package dto

import (
	"github.com/TanPhat-26/Go-ecommerce-backend-api/internal/models"
	"github.com/google/uuid"
)

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Phone     string    `json:"phone"`
	Status    string    `json:"status"`
}

// ToUserResponse maps the internal User model to a public API response
// Sensitive and internal field such as PasswordHash and DeletedAt
// are intentionally excluded from the response
func ToUserResponse(user *models.User) UserResponse {
	response := UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Status:    user.Status,
	}
	if user.Phone != nil {
		response.Phone = *user.Phone
	}
	return response
}
