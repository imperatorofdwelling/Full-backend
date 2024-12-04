package staysreports

import (
	"errors"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	handler "github.com/imperatorofdwelling/Full-backend/internal/api/handler/user"
	"github.com/imperatorofdwelling/Full-backend/internal/config"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces/mocks"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestStaysReportsHandler_NewStaysReportsHandler(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.StaysReportsService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	t.Run("should be no errors", func(t *testing.T) {
		hdl.NewStaysReportsHandler(router)
	})
}

func TestStaysReportsHandler_UserIdError(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.StaysReportsService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	router.Get("/report", hdl.GetAllStaysReports)
	router.Post("/report/create", hdl.CreateStaysReports)
	router.Put("/report/{reportId}", hdl.UpdateStaysReports)
	router.Delete("/report/{reportId}", hdl.DeleteStaysReports)

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("get user error", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/report", nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)
		assert.Equal(t, http.StatusUnauthorized, r.Code)
	})
	t.Run("post user error", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPost, "/report/create", nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)
		assert.Equal(t, http.StatusUnauthorized, r.Code)
	})
	t.Run("put user error", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPut, "/report/"+testUserID.String(), nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)
		assert.Equal(t, http.StatusUnauthorized, r.Code)
	})
	t.Run("delete user error", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodDelete, "/report/"+testUserID.String(), nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)
		assert.Equal(t, http.StatusUnauthorized, r.Code)
	})
}

func TestStaysReportsHandler_ParamsError(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.StaysReportsService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	router.Use(handler.JWTMiddleware("your-secret-key", log))

	router.Post("/report/create/{stayId}", hdl.CreateStaysReports)
	router.Delete("/report/{reportId}", hdl.DeleteStaysReports)
	router.Put("/report/{reportId}", hdl.UpdateStaysReports)
	router.Get("/report", hdl.GetAllStaysReports)

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("should be delete error", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodDelete, "/report/"+testUserID.String(), nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("DeleteStaysReports", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("failed to delete stay report"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
	t.Run("should be get error", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/report", nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("GetAllStaysReports", mock.Anything, mock.Anything).Return(nil, errors.New("failed to fetch reports"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
	t.Run("should be params errors put", func(t *testing.T) {
		r := httptest.NewRecorder()

		payload := `{"title": "smth"}`

		req := httptest.NewRequest(http.MethodPut, "/report/"+testUserID.String(), strings.NewReader(payload))
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
	t.Run("should be body params empty errors put", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPut, "/report/"+testUserID.String(), nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
	t.Run("should be service error put", func(t *testing.T) {
		r := httptest.NewRecorder()

		payload := `{"title": "smth", "description": "smth"}`

		req := httptest.NewRequest(http.MethodPut, "/report/"+testUserID.String(), strings.NewReader(payload))
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("UpdateStaysReports", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("failed to create stays report"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
	t.Run("should be params errors post", func(t *testing.T) {
		r := httptest.NewRecorder()

		payload := `{"title": "smth"}`

		req := httptest.NewRequest(http.MethodPost, "/report/create/"+testUserID.String(), strings.NewReader(payload))
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
	t.Run("should be body params empty errors post", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPost, "/report/create/"+testUserID.String(), nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
	t.Run("should be service error post", func(t *testing.T) {
		r := httptest.NewRecorder()

		payload := `{"title": "smth", "description": "smth"}`

		req := httptest.NewRequest(http.MethodPost, "/report/create/"+testUserID.String(), strings.NewReader(payload))
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("CreateStaysReports", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("failed to create stays report"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})

}

func TestStaysReportsHandler_ReportCreateSuccess(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.StaysReportsService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()
	router.Use(handler.JWTMiddleware("your-secret-key", log))
	router.Post("/report/create/{stayId}", hdl.CreateStaysReports)
	router.Delete("/report/{reportId}", hdl.DeleteStaysReports)
	router.Put("/report/{reportId}", hdl.UpdateStaysReports)
	router.Get("/report", hdl.GetAllStaysReports)

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("should be delete success", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodDelete, "/report/"+testUserID.String(), nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("DeleteStaysReports", mock.Anything, mock.Anything, mock.Anything).Return(nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})
	t.Run("should be get success", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/report", nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("GetAllStaysReports", mock.Anything, mock.Anything).Return(nil, nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)

	})
	t.Run("should be success creating a new report", func(t *testing.T) {
		r := httptest.NewRecorder()

		payload := `{"title": "smth", "description": "smth"}`

		req := httptest.NewRequest(http.MethodPost, "/report/create/"+testUserID.String(), strings.NewReader(payload))
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("CreateStaysReports", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusCreated, r.Code)
	})
	t.Run("should be success updating a new report", func(t *testing.T) {
		r := httptest.NewRecorder()

		payload := `{"title": "smth", "description": "smth"}`

		req := httptest.NewRequest(http.MethodPut, "/report/"+testUserID.String(), strings.NewReader(payload))
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("UpdateStaysReports", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})
}
