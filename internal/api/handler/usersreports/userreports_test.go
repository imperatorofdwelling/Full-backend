package usersreports

import (
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	handler "github.com/imperatorofdwelling/Full-backend/internal/api/handler/user"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces/mocks"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestUsersReportsHandler_NewUsersReportsHandler(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.UsersReportsService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	t.Run("should be no errors", func(t *testing.T) {
		hdl.NewUsersReportsHandler(router)
	})
}

func TestUsersReportsHandler_UserIdError(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.UsersReportsService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()
	router.Get("/user/report", hdl.GetAllUsersReports)
	router.Post("/user/report/create/{toBlameId}", hdl.CreateUsersReports)
	router.Put("/user/report/{reportId}", hdl.UpdateUsersReports)
	router.Delete("/user/report/{reportId}", hdl.DeleteUsersReports)

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("get user error", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/user/report", nil)
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

		req := httptest.NewRequest(http.MethodPost, "/user/report/create/"+testUserID.String(), nil)
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

		req := httptest.NewRequest(http.MethodPut, "/user/report/"+testUserID.String(), nil)
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

		req := httptest.NewRequest(http.MethodDelete, "/user/report/"+testUserID.String(), nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusUnauthorized, r.Code)
	})
}

func TestUsersReportsHandler_ParamsError(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.UsersReportsService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	router.Use(handler.JWTMiddleware("your-secret-key", log))

	router.Get("/user/report", hdl.GetAllUsersReports)
	router.Post("/user/report/create/{toBlameId}", hdl.CreateUsersReports)
	router.Put("/user/report/{reportId}", hdl.UpdateUsersReports)
	router.Delete("/user/report/{reportId}", hdl.DeleteUsersReports)

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("should be params errors post", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPost, "/user/report/create/"+testUserID.String(), nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("CreateUsersReports", mock.Anything, mock.Anything).Return(nil, errors.New("failed to fetch reports"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
	t.Run("should be params errors put", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPut, "/user/report/"+testUserID.String(), nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("UpdateUsersReports", mock.Anything, mock.Anything).Return(nil, errors.New("failed to fetch reports"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
	t.Run("should be params errors put", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodDelete, "/user/report/"+testUserID.String(), nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("DeleteUsersReports", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("failed to fetch reports"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
	t.Run("should be not full params errors post", func(t *testing.T) {
		r := httptest.NewRecorder()

		payload := `{"title": "smth"}`

		req := httptest.NewRequest(http.MethodPost, "/user/report/create/"+testUserID.String(), strings.NewReader(payload))
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("CreateUsersReports", mock.Anything, mock.Anything).Return(nil, errors.New("failed to fetch reports"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
	t.Run("should be not full params errors put", func(t *testing.T) {
		r := httptest.NewRecorder()

		payload := `{"title": "smth"}`

		req := httptest.NewRequest(http.MethodPut, "/user/report/"+testUserID.String(), strings.NewReader(payload))
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("UpdateUsersReports", mock.Anything, mock.Anything).Return(nil, errors.New("failed to fetch reports"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
	t.Run("should be body params empty errors post", func(t *testing.T) {
		r := httptest.NewRecorder()

		payload := `{"title": "smth", "description": "smth"}`

		req := httptest.NewRequest(http.MethodPost, "/user/report/create/"+testUserID.String(), strings.NewReader(payload))
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("CreateUsersReports", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("failed to create stays report"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
	t.Run("should be body params empty errors put", func(t *testing.T) {
		r := httptest.NewRecorder()

		payload := `{"title": "smth", "description": "smth"}`

		req := httptest.NewRequest(http.MethodPut, "/user/report/"+testUserID.String(), strings.NewReader(payload))
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("UpdateUsersReports", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("failed to update stays report"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})

	t.Run("should be get service error", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/user/report", nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("GetAllUsersReports", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("failed to update stays report"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestUsersReportsHandler_Success(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.UsersReportsService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	router.Use(handler.JWTMiddleware("your-secret-key", log))

	router.Get("/user/report", hdl.GetAllUsersReports)
	router.Post("/user/report/create/{toBlameId}", hdl.CreateUsersReports)
	router.Put("/user/report/{reportId}", hdl.UpdateUsersReports)
	router.Delete("/user/report/{reportId}", hdl.DeleteUsersReports)

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("should be delete success", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodDelete, "/user/report/"+testUserID.String(), nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("DeleteUsersReports", mock.Anything, mock.Anything, mock.Anything).Return(nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})
	t.Run("should be get success", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/user/report", nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("GetAllUsersReports", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})
	t.Run("should be put success", func(t *testing.T) {
		r := httptest.NewRecorder()

		payload := `{"title": "smth", "description": "smth"}`

		req := httptest.NewRequest(http.MethodPut, "/user/report/"+testUserID.String(), strings.NewReader(payload))
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("UpdateUsersReports", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})
	t.Run("should be post success", func(t *testing.T) {
		r := httptest.NewRecorder()

		payload := `{"title": "smth", "description": "smth"}`

		req := httptest.NewRequest(http.MethodPost, "/user/report/create/"+testUserID.String(), strings.NewReader(payload))
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("CreateUsersReports", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusCreated, r.Code)
	})
}
