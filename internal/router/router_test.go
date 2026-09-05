package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TanPhat-26/Go-ecommerce-backend-api/global"
	"github.com/TanPhat-26/Go-ecommerce-backend-api/pkg/setting"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func TestHealthRoute(t *testing.T) {
	global.Config = &setting.Config{
		App: setting.APPSetting{
			Name: "go-ecommerce-backend-api",
			Env:  "test",
			Port: "8080",
		},
	}

	r := NewRouter()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	if w.Header().Get("X-Request-ID") == "" {
		t.Fatal("expected X-Request-ID header")
	}
}

func TestRecoveryMiddleware(t *testing.T) {
	global.Config = &setting.Config{
		App: setting.APPSetting{
			Name: "go-ecommerce-backend-api",
			Env:  "test",
			Port: "8080",
		},
	}

	global.Logger = zap.NewNop()

	r := NewRouter()
	r.GET("/panic", func(c *gin.Context) {
		panic("test panic")
	})

	req := httptest.NewRequest(http.MethodGet, "/panic", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected status %d, got %d", http.StatusInternalServerError, w.Code)
	}
}
