package dto

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestRegisterRequestValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name    string
		request RegisterRequest
		wantErr bool
	}{
		{
			name: "valid request",
			request: RegisterRequest{
				Email:     "user@example.com",
				Password:  "password123",
				FirstName: "Tan",
				LastName:  "Phat",
			},
			wantErr: false,
		},
		{
			name: "invalid email",
			request: RegisterRequest{
				Email:     "invalid-email",
				Password:  "password123",
				FirstName: "Tan",
				LastName:  "Phat",
			},
			wantErr: true,
		},
		{
			name: "short password",
			request: RegisterRequest{
				Email:     "user@example.com",
				Password:  "short",
				FirstName: "Tan",
				LastName:  "Phat",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.Struct(tt.request)
			if (err != nil) != tt.wantErr {
				t.Fatalf("validation error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
