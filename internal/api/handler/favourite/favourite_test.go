package favourite

import (
	"errors"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	handler "github.com/imperatorofdwelling/Full-backend/internal/api/handler/user"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces/mocks"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestFavHandler_NewFavouriteHandler(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.FavouriteService{}
	hdl := FavHandler{
		Log: log,
		Svc: &svc,
	}

	router := chi.NewRouter()

	t.Run("should be no error", func(t *testing.T) {
		hdl.NewFavouriteHandler(router)
	})
}

func TestFavHandler_AddFavouriteUserID(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.FavouriteService{}
	hdl := FavHandler{
		Log: log,
		Svc: &svc,
	}

	router := chi.NewRouter()

	router.Post("/favourites/{stayID}", hdl.AddFavourite)

	testUserID, _ := uuid.NewV4()

	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("should be error with the user id", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPost, "/favourites/"+testUserID.String(), nil)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)
		assert.Equal(t, http.StatusUnauthorized, r.Code)
	})

}

func TestFavHandler_AddFavouriteSuccess(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.FavouriteService{}
	hdl := FavHandler{
		Log: log,
		Svc: &svc,
	}

	router := chi.NewRouter()
	router.Use(handler.JWTMiddleware("your-secret-key", log))
	router.Post("/favourites/{stayID}", hdl.AddFavourite)

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("should be error with the stay id", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("AddToFavourites", mock.Anything, testUserID.String(), mock.Anything).Return(nil)

		req := httptest.NewRequest(http.MethodPost, "/favourites/someStayID", nil)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)
		assert.Equal(t, http.StatusOK, r.Code)
	})
}

func TestFavHandler_AddFavouriteStayId(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.FavouriteService{}
	hdl := FavHandler{
		Log: log,
		Svc: &svc,
	}

	router := chi.NewRouter()

	router.Use(handler.JWTMiddleware("your-secret-key", log))
	router.Post("/favourites/{stayID}", hdl.AddFavourite)

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID,
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("should be error with the stay id", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("AddToFavourites", mock.Anything, testUserID.String(), mock.Anything).Return(errors.New("stay id is wrong"))

		req := httptest.NewRequest(http.MethodPost, "/favourites/someStayID", nil)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)
		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestFavHandler_GetAllFavouritesUserId(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.FavouriteService{}
	hdl := FavHandler{
		Log: log,
		Svc: &svc,
	}

	router := chi.NewRouter()

	router.Get("/favourites", hdl.GetAllFavourites)

	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("should return unauthorized when user id is missing", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/favourites", nil)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusUnauthorized, r.Code)
	})
}

func TestFavHandler_GetAllFavourites(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.FavouriteService{}
	hdl := FavHandler{
		Log: log,
		Svc: &svc,
	}

	router := chi.NewRouter()
	router.Use(handler.JWTMiddleware("your-secret-key", log))
	router.Get("/favourites", hdl.GetAllFavourites)

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("should return error when service method fails", func(t *testing.T) {
		r := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/favourites", nil)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("GetAllFavourites", mock.Anything, testUserID.String()).Return(nil, errors.New("service error"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestFavHandler_GetAllFavouritesSuccess(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.FavouriteService{}
	hdl := FavHandler{
		Log: log,
		Svc: &svc,
	}

	router := chi.NewRouter()
	router.Use(handler.JWTMiddleware("your-secret-key", log))
	router.Get("/favourites", hdl.GetAllFavourites)

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("should return error when service method fails", func(t *testing.T) {
		r := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/favourites", nil)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("GetAllFavourites", mock.Anything, testUserID.String()).Return(nil, nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})
}

func TestFavHandler_RemoveFavouriteUserId(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.FavouriteService{}
	hdl := FavHandler{
		Log: log,
		Svc: &svc,
	}

	router := chi.NewRouter()

	router.Delete("/favourites/{stayID}", hdl.RemoveFavourite)

	testUserID, _ := uuid.NewV4()

	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("should be error with the user id", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodDelete, "/favourites/"+testUserID.String(), nil)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)
		assert.Equal(t, http.StatusUnauthorized, r.Code)
	})

}

func TestFavHandler_RemoveFavouriteStayId(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.FavouriteService{}
	hdl := FavHandler{
		Log: log,
		Svc: &svc,
	}

	router := chi.NewRouter()

	router.Use(handler.JWTMiddleware("your-secret-key", log))
	router.Delete("/favourites/{stayID}", hdl.RemoveFavourite)

	testUserID, _ := uuid.NewV4()

	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("should be error with the user id", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("RemoveFromFavourites", mock.Anything, testUserID.String(), mock.Anything).Return(errors.New("stay id is incorrect"))
		req := httptest.NewRequest(http.MethodDelete, "/favourites/"+testUserID.String(), nil)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)
		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})

}

func TestFavHandler_RemoveFavouriteSuccess(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.FavouriteService{}
	hdl := FavHandler{
		Log: log,
		Svc: &svc,
	}

	router := chi.NewRouter()

	router.Use(handler.JWTMiddleware("your-secret-key", log))
	router.Delete("/favourites/{stayID}", hdl.RemoveFavourite)

	testUserID, _ := uuid.NewV4()

	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("should be error with the user id", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("RemoveFromFavourites", mock.Anything, testUserID.String(), mock.Anything).Return(nil)
		req := httptest.NewRequest(http.MethodDelete, "/favourites/"+testUserID.String(), nil)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)
		assert.Equal(t, http.StatusOK, r.Code)
	})

}
