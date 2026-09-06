package service

import (
	"context"
	"errors"
	"testing"

	"github.com/TanPhat-26/Go-ecommerce-backend-api/internal/dto"
	"github.com/TanPhat-26/Go-ecommerce-backend-api/internal/models"
	"github.com/TanPhat-26/Go-ecommerce-backend-api/internal/repo"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type fakeUserRepository struct {
	findByEmailFn func(context.Context, string) (*models.User, error)
	lookedUpEmail string
}

func (f *fakeUserRepository) Create(context.Context, *models.User) error {
	return nil
}

func (f *fakeUserRepository) FindByID(context.Context, uuid.UUID) (*models.User, error) {
	return nil, errors.New("not implemented")
}

func (f *fakeUserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	f.lookedUpEmail = email
	return f.findByEmailFn(ctx, email)
}

func (f *fakeUserRepository) Update(context.Context, *models.User) error {
	return nil
}

type fakeRegistrationRepository struct {
	createdUser *models.User
	roleName    string
	err         error
}

func (f *fakeRegistrationRepository) CreateUserWithRole(
	_ context.Context,
	user *models.User,
	roleName string,
) error {
	f.createdUser = user
	f.roleName = roleName
	return f.err
}

func TestAuthServiceRegister(t *testing.T) {
	request := dto.RegisterRequest{
		Email:     " User@Example.com ",
		Password:  "password123",
		FirstName: " Tan ",
		LastName:  " Phat ",
		Phone:     " 0900000000 ",
	}

	t.Run("registers a normalized customer account with a bcrypt password", func(t *testing.T) {
		userRepository := &fakeUserRepository{
			findByEmailFn: func(context.Context, string) (*models.User, error) {
				return nil, gorm.ErrRecordNotFound
			},
		}
		registrationRepository := &fakeRegistrationRepository{}
		authService := NewAuthService(userRepository, registrationRepository)

		response, err := authService.Register(context.Background(), request)
		if err != nil {
			t.Fatalf("Register() error = %v", err)
		}

		if userRepository.lookedUpEmail != "user@example.com" {
			t.Fatalf("lookup email = %q, want normalized email", userRepository.lookedUpEmail)
		}
		if registrationRepository.roleName != "customer" {
			t.Fatalf("role name = %q, want customer", registrationRepository.roleName)
		}
		if registrationRepository.createdUser == nil {
			t.Fatal("expected registration repository to receive a user")
		}
		if registrationRepository.createdUser.PasswordHash == request.Password {
			t.Fatal("plaintext password must not be stored")
		}
		if err := bcrypt.CompareHashAndPassword(
			[]byte(registrationRepository.createdUser.PasswordHash),
			[]byte(request.Password),
		); err != nil {
			t.Fatalf("password hash is invalid: %v", err)
		}
		if response.Email != "user@example.com" || response.FirstName != "Tan" || response.LastName != "Phat" {
			t.Fatalf("unexpected response: %+v", response)
		}
		if response.Phone != "0900000000" || response.Status != "active" {
			t.Fatalf("unexpected response: %+v", response)
		}
	})

	t.Run("returns a conflict error when the email already exists", func(t *testing.T) {
		userRepository := &fakeUserRepository{
			findByEmailFn: func(context.Context, string) (*models.User, error) {
				return &models.User{}, nil
			},
		}
		registrationRepository := &fakeRegistrationRepository{}
		authService := NewAuthService(userRepository, registrationRepository)

		_, err := authService.Register(context.Background(), request)
		if !errors.Is(err, ErrEmailAlreadyExists) {
			t.Fatalf("Register() error = %v, want ErrEmailAlreadyExists", err)
		}
		if registrationRepository.createdUser != nil {
			t.Fatal("registration must not run when the email exists")
		}
	})

	t.Run("maps a database uniqueness error to the public conflict error", func(t *testing.T) {
		userRepository := &fakeUserRepository{
			findByEmailFn: func(context.Context, string) (*models.User, error) {
				return nil, gorm.ErrRecordNotFound
			},
		}
		registrationRepository := &fakeRegistrationRepository{err: repo.ErrEmailAlreadyExists}
		authService := NewAuthService(userRepository, registrationRepository)

		_, err := authService.Register(context.Background(), request)
		if !errors.Is(err, ErrEmailAlreadyExists) {
			t.Fatalf("Register() error = %v, want ErrEmailAlreadyExists", err)
		}
	})

	t.Run("wraps an unexpected lookup error", func(t *testing.T) {
		lookupError := errors.New("database unavailable")
		userRepository := &fakeUserRepository{
			findByEmailFn: func(context.Context, string) (*models.User, error) {
				return nil, lookupError
			},
		}
		authService := NewAuthService(userRepository, &fakeRegistrationRepository{})

		_, err := authService.Register(context.Background(), request)
		if !errors.Is(err, lookupError) {
			t.Fatalf("Register() error = %v, want wrapped lookup error", err)
		}
	})
}
