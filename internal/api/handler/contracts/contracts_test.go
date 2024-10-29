package contracts

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

func TestContracts_NewContractsHandler(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.ContractService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	t.Run("should be no errors", func(t *testing.T) {
		hdl.NewContractHandler(router)
	})
}

func TestContracts_ContractGetAllContactsUserId(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.ContractService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	router := chi.NewRouter()
	router.Get("/contract", hdl.GetAllContracts)

	t.Run("should be user id error", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/contract", nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)
		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusUnauthorized, r.Code)
	})
}

func TestContracts_ContractGetAllContractsError(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.ContractService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	router := chi.NewRouter()
	router.Use(handler.JWTMiddleware("your-secret-key", log))
	router.Get("/contract", hdl.GetAllContracts)

	t.Run("should be user id error", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/contract", nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("GetAllContracts", mock.Anything, testUserID.String()).Return(nil, errors.New("could not fetch contracts"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestContracts_ContractGetAllContractsSuccess(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.ContractService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	router := chi.NewRouter()
	router.Use(handler.JWTMiddleware("your-secret-key", log))
	router.Get("/contract", hdl.GetAllContracts)

	t.Run("should be user id error", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/contract", nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("GetAllContracts", mock.Anything, testUserID.String()).Return(nil, nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})
}

func TestContracts_ContractAddContractUserIdError(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.ContractService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	router := chi.NewRouter()
	router.Post("/contract", hdl.AddContract)
	router.Put("/contract", hdl.UpdateContract)

	t.Run("should be user id error post", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPost, "/contract", nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)
		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusUnauthorized, r.Code)
	})

	t.Run("should be user id error put", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPut, "/contract", nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)
		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusUnauthorized, r.Code)
	})
}

func TestContracts_ContractAddContractMapError(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.ContractService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	router := chi.NewRouter()
	router.Use(handler.JWTMiddleware("your-secret-key", log))
	router.Post("/contract", hdl.AddContract)
	router.Put("/contract", hdl.UpdateContract)

	t.Run("should return bad request error on decode failure post", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPost, "/contract", strings.NewReader("{invalid-json"))
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should return bad request error on decode failure put", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPut, "/contract", strings.NewReader("{invalid-json"))
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
}

func TestContracts_ContractAddContractMissingFields(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.ContractService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	router := chi.NewRouter()
	router.Use(handler.JWTMiddleware("your-secret-key", log))
	router.Post("/contract", hdl.AddContract)
	router.Put("/contract", hdl.UpdateContract)

	t.Run("should return bad request error on missing dateStart or dateEnd post", func(t *testing.T) {
		r := httptest.NewRecorder()

		payload := `{"dateStart": "2023-10-30T10:00:00Z"}`
		req := httptest.NewRequest(http.MethodPost, "/contract", strings.NewReader(payload))
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should return bad request error on missing dateStart or dateEnd put", func(t *testing.T) {
		r := httptest.NewRecorder()

		payload := `{"dateStart": "2023-10-30T10:00:00Z"}`
		req := httptest.NewRequest(http.MethodPut, "/contract", strings.NewReader(payload))
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
}

func TestContracts_ContractAddContractInvalidDateFormat(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.ContractService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	router := chi.NewRouter()
	router.Use(handler.JWTMiddleware("your-secret-key", log))
	router.Post("/contract", hdl.AddContract)
	router.Put("/contract", hdl.UpdateContract)

	t.Run("should return bad request error on invalid dateStart format post", func(t *testing.T) {
		r := httptest.NewRecorder()

		payload := `{"dateStart": "invalid-date", "dateEnd": "2023-11-01T10:00:00Z"}`
		req := httptest.NewRequest(http.MethodPost, "/contract", strings.NewReader(payload))
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should return bad request error on invalid dateEnd format post", func(t *testing.T) {
		r := httptest.NewRecorder()

		payload := `{"dateStart": "2023-10-30T10:00:00Z", "dateEnd": "invalid-date"}`
		req := httptest.NewRequest(http.MethodPost, "/contract", strings.NewReader(payload))
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should return bad request error on invalid dateStart format put", func(t *testing.T) {
		r := httptest.NewRecorder()

		payload := `{"dateStart": "invalid-date", "dateEnd": "2023-11-01T10:00:00Z"}`
		req := httptest.NewRequest(http.MethodPut, "/contract", strings.NewReader(payload))
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should return bad request error on invalid dateEnd format put", func(t *testing.T) {
		r := httptest.NewRecorder()

		payload := `{"dateStart": "2023-10-30T10:00:00Z", "dateEnd": "invalid-date"}`
		req := httptest.NewRequest(http.MethodPut, "/contract", strings.NewReader(payload))
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
}

func TestContracts_AddContract_Error(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.ContractService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	router := chi.NewRouter()
	router.Use(handler.JWTMiddleware("your-secret-key", log))
	router.Post("/contract", hdl.AddContract)
	router.Put("/contract", hdl.UpdateContract)

	t.Run("should be add contract error post", func(t *testing.T) {
		r := httptest.NewRecorder()

		payload := `{"dateStart": "2023-11-01T10:00:00Z", "dateEnd": "2023-11-01T10:00:00Z"}`
		req := httptest.NewRequest(http.MethodPost, "/contract", strings.NewReader(payload))
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("AddContract", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(errors.New("could not add contract"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})

	t.Run("should be add contract error put", func(t *testing.T) {
		r := httptest.NewRecorder()

		payload := `{"dateStart": "2023-11-01T10:00:00Z", "dateEnd": "2023-11-01T10:00:00Z"}`
		req := httptest.NewRequest(http.MethodPut, "/contract", strings.NewReader(payload))
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("UpdateContract", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(errors.New("could not add contract"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestContracts_AddContract_Success(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.ContractService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	router := chi.NewRouter()
	router.Use(handler.JWTMiddleware("your-secret-key", log))
	router.Post("/contract", hdl.AddContract)
	router.Put("/contract", hdl.UpdateContract)

	t.Run("should be add contract error post", func(t *testing.T) {
		r := httptest.NewRecorder()

		payload := `{"dateStart": "2023-11-01T10:00:00Z", "dateEnd": "2023-11-01T10:00:00Z"}`
		req := httptest.NewRequest(http.MethodPost, "/contract", strings.NewReader(payload))
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("AddContract", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusCreated, r.Code)
	})

	t.Run("should be add contract error put", func(t *testing.T) {
		r := httptest.NewRecorder()

		payload := `{"dateStart": "2023-11-01T10:00:00Z", "dateEnd": "2023-11-01T10:00:00Z"}`
		req := httptest.NewRequest(http.MethodPut, "/contract", strings.NewReader(payload))
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("UpdateContract", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusCreated, r.Code)
	})
}
