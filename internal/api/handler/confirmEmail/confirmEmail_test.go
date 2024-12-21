package confirmEmail

import (
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/config"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces/mocks"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler_NewConfirmEmailHandler(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	router := chi.NewRouter()

	log := logger.New()
	svc := mocks.ConfirmEmailService{}

	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	t.Run("should be no errors", func(t *testing.T) {
		hdl.NewConfirmEmailHandler(router)
	})
}

func TestHandler_ConfirmEmail_CreateOTP_UserID_Error(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.ConfirmEmailService{}

	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()
	router.Get("/otp", hdl.CreateOTP)

	t.Run("should be user id error", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/otp", nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusUnauthorized, r.Code)
	})
}

func TestHandler_ConfirmEmail_CreateOTP_SVC_Success(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.ConfirmEmailService{}

	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()
	router.Get("/otp", hdl.CreateOTP)

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("should be no user id error", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/otp", nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		ctx := context.WithValue(req.Context(), "user_id", testUserID.String())
		req = req.WithContext(ctx)

		svc.On("CreateOTP", mock.Anything, mock.Anything).Return(nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})
}

func TestHandler_ConfirmEmail_CreateOTP_SVC_Error(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.ConfirmEmailService{}

	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()
	router.Get("/otp", hdl.CreateOTP)

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("should be no user id error", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/otp", nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		ctx := context.WithValue(req.Context(), "user_id", testUserID.String())
		req = req.WithContext(ctx)

		svc.On("CreateOTP", mock.Anything, mock.Anything).Return(errors.New("error creating otp"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}
