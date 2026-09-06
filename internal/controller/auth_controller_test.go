package controller

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TanPhat-26/Go-ecommerce-backend-api/internal/dto"
	"github.com/TanPhat-26/Go-ecommerce-backend-api/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type fakeAuthService struct {
	registerFn func(context.Context, dto.RegisterRequest) (*dto.UserResponse, error)
	called     bool
}

func (f *fakeAuthService) Register(
	ctx context.Context,
	request dto.RegisterRequest,
) (*dto.UserResponse, error) {
	f.called = true
	return f.registerFn(ctx, request)
}

func TestAuthControllerRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("returns created for a valid registration", func(t *testing.T) {
		fakeService := &fakeAuthService{
			registerFn: func(context.Context, dto.RegisterRequest) (*dto.UserResponse, error) {
				return &dto.UserResponse{
					ID:        uuid.New(),
					Email:     "user@example.com",
					FirstName: "Tan",
					LastName:  "Phat",
					Status:    "active",
				}, nil
			},
		}
		controller := NewAuthController(fakeService)
		router := gin.New()
		router.POST("/register", controller.Register)

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodPost,
			"/register",
			bytes.NewBufferString(`{"email":"user@example.com","password":"password123","first_name":"Tan","last_name":"Phat"}`),
		)
		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(recorder, request)

		if recorder.Code != http.StatusCreated {
			t.Fatalf("status = %d, want %d", recorder.Code, http.StatusCreated)
		}
		if !fakeService.called {
			t.Fatal("expected service to be called")
		}
	})

	t.Run("rejects invalid input before calling the service", func(t *testing.T) {
		fakeService := &fakeAuthService{
			registerFn: func(context.Context, dto.RegisterRequest) (*dto.UserResponse, error) {
				return nil, nil
			},
		}
		controller := NewAuthController(fakeService)
		router := gin.New()
		router.POST("/register", controller.Register)

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodPost,
			"/register",
			bytes.NewBufferString(`{"email":"invalid","password":"short"}`),
		)
		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(recorder, request)

		if recorder.Code != http.StatusBadRequest {
			t.Fatalf("status = %d, want %d", recorder.Code, http.StatusBadRequest)
		}
		if fakeService.called {
			t.Fatal("service must not be called for invalid input")
		}
	})

	t.Run("returns conflict when the email is already registered", func(t *testing.T) {
		fakeService := &fakeAuthService{
			registerFn: func(context.Context, dto.RegisterRequest) (*dto.UserResponse, error) {
				return nil, service.ErrEmailAlreadyExists
			},
		}
		controller := NewAuthController(fakeService)
		router := gin.New()
		router.POST("/register", controller.Register)

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodPost,
			"/register",
			bytes.NewBufferString(`{"email":"user@example.com","password":"password123","first_name":"Tan","last_name":"Phat"}`),
		)
		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(recorder, request)

		if recorder.Code != http.StatusConflict {
			t.Fatalf("status = %d, want %d", recorder.Code, http.StatusConflict)
		}
	})
}
