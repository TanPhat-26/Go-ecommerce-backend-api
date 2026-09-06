package controller

import (
	"errors"
	"net/http"

	"github.com/TanPhat-26/Go-ecommerce-backend-api/internal/dto"
	"github.com/TanPhat-26/Go-ecommerce-backend-api/internal/service"
	"github.com/TanPhat-26/Go-ecommerce-backend-api/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// AuthController handles HTTP requests for authentication workflows
type AuthController struct {
	authService service.AuthService
	validator   *validator.Validate
}

// NewAuthController creates an authentication controller with its service dependency
func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
		validator:   utils.NewValidator(),
	}
}

// Register validates a registration request and creates a customer account
func (a *AuthController) Register(c *gin.Context) {
	var request dto.RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "invalid request body",
		})
		return
	}

	if err := a.validator.Struct(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "validation failed",
		})
		return
	}

	user, err := a.authService.Register(c.Request.Context(), request)
	if err != nil {
		if errors.Is(err, service.ErrEmailAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{
				"status":  "error",
				"message": "email already exists",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "could not register user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "user registered successfully",
		"data":    user,
	})
}
