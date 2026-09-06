package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/TanPhat-26/Go-ecommerce-backend-api/internal/dto"
	"github.com/TanPhat-26/Go-ecommerce-backend-api/internal/models"
	"github.com/TanPhat-26/Go-ecommerce-backend-api/internal/repo"
	"gorm.io/gorm"
)

const defaultCustomerRoleName = "customer"

// AuthService defines authentication-related business operations
// The interface keeps controllers independent from the concrete service implementation
type AuthService interface {
	Register(
		ctx context.Context,
		req dto.RegisterRequest,
	) (*dto.UserResponse, error)
}

type authService struct {
	userRepo         repo.UserRepository
	registrationRepo repo.RegistrationRepository
}

// NewAuthService creates an authentication service with its persistence dependencies
func NewAuthService(
	userRepo repo.UserRepository,
	registrationRepo repo.RegistrationRepository,
) AuthService {
	return &authService{
		userRepo:         userRepo,
		registrationRepo: registrationRepo,
	}
}

// ErrEmailAlreadyExists indicates that the email is already registered
var ErrEmailAlreadyExists = errors.New("email already exists")

func (s *authService) Register(
	ctx context.Context,
	req dto.RegisterRequest,
) (*dto.UserResponse, error) {
	// Normalize email before lookup to prevent duplicate accounts caused by casing or whitespace
	email := strings.ToLower(strings.TrimSpace(req.Email))

	existingUser, err := s.userRepo.FindByEmail(ctx, email)
	if err == nil && existingUser != nil {
		return nil, ErrEmailAlreadyExists
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("check email uniqueness: %w", err)
	}

	// Password hashing belongs to the service layer because it is part of
	// the registration business rule, not a database persistence concern
	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}
	user := &models.User{
		Email:        email,
		PasswordHash: string(passwordHash),
		FirstName:    strings.TrimSpace(req.FirstName),
		LastName:     strings.TrimSpace(req.LastName),
		Status:       "active",
	}

	phone := strings.TrimSpace(req.Phone)
	if phone != "" {
		user.Phone = &phone
	}

	// Create the user and assign the default customer role atomically.
	if err := s.registrationRepo.CreateUserWithRole(
		ctx,
		user,
		defaultCustomerRoleName,
	); err != nil {
		if errors.Is(err, repo.ErrEmailAlreadyExists) {
			return nil, ErrEmailAlreadyExists
		}

		return nil, fmt.Errorf("register user: %w", err)
	}
	resp := dto.ToUserResponse(user)
	return &resp, nil
}
