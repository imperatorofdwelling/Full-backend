package staysreviews

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces/mocks"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/staysreviews"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestStaysReviewsHandler_NewStaysReviewsHandler(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.StaysReviewsService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	t.Run("should be no errors", func(t *testing.T) {
		hdl.NewStaysReviewsHandler(router)
	})
}

func TestStaysReviewsHandler_CreateStaysReviewHandler(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.StaysReviewsService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	t.Run("should be no errors", func(t *testing.T) {
		r := httptest.NewRecorder()

		uuidStay, _ := uuid.NewV4()

		uuidUser, _ := uuid.NewV4()

		payload := staysreviews.StaysReviewEntity{
			StayID:      uuidStay,
			UserID:      uuidUser,
			Title:       "test",
			Description: "test",
			Rating:      1.2,
		}

		pBytes, _ := json.Marshal(payload)

		pBuf := bytes.NewBuffer(pBytes)

		svc.On("CreateStaysReview", mock.Anything, mock.Anything).Return(nil).Once()

		req := httptest.NewRequest(http.MethodPost, "/staysreviews/create", pBuf)

		router.HandleFunc("/staysreviews/create", hdl.CreateStaysReview)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusCreated, r.Code)
	})

	t.Run("should be error decoding body", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPost, "/staysreviews/create", strings.NewReader(""))

		router.HandleFunc("/staysreviews/create", hdl.CreateStaysReview)

		router.ServeHTTP(r, req)
		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should be error creating stays review", func(t *testing.T) {
		r := httptest.NewRecorder()

		uuidStay, _ := uuid.NewV4()

		uuidUser, _ := uuid.NewV4()

		payload := staysreviews.StaysReviewEntity{
			StayID:      uuidStay,
			UserID:      uuidUser,
			Title:       "test",
			Description: "test",
			Rating:      1.2,
		}

		pBytes, _ := json.Marshal(payload)

		pBuf := bytes.NewBuffer(pBytes)

		svc.On("CreateStaysReview", mock.Anything, mock.Anything).Return(errors.New("error creating stays review")).Once()

		req := httptest.NewRequest(http.MethodPost, "/staysreviews/create", pBuf)

		router.HandleFunc("/staysreviews/create", hdl.CreateStaysReview)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestStaysReviewsHandler_UpdateStaysReviewHandler(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.StaysReviewsService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	t.Run("should be no errors", func(t *testing.T) {
		r := httptest.NewRecorder()

		uuidStayReview, _ := uuid.NewV4()

		uuidStay, _ := uuid.NewV4()

		uuidUser, _ := uuid.NewV4()

		payload := staysreviews.StaysReviewEntity{
			StayID:      uuidStay,
			UserID:      uuidUser,
			Title:       "test",
			Description: "test",
			Rating:      1.2,
		}

		pBytes, _ := json.Marshal(payload)

		pBuf := bytes.NewBuffer(pBytes)

		svc.On("UpdateStaysReview", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil).Once()

		req := httptest.NewRequest(http.MethodPut, "/staysreviews/update/"+uuidStayReview.String(), pBuf)

		router.HandleFunc("/staysreviews/update/{id}", hdl.UpdateStaysReview)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})

	t.Run("should be error parsing uuid", func(t *testing.T) {
		r := httptest.NewRecorder()

		fakeID := "fake"

		req := httptest.NewRequest(http.MethodPut, "/staysreviews/update/"+fakeID, nil)

		router.HandleFunc("/staysreviews/update/{id}", hdl.UpdateStaysReview)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should be error parsing body", func(t *testing.T) {
		r := httptest.NewRecorder()

		uuidStayReview, _ := uuid.NewV4()

		req := httptest.NewRequest(http.MethodPut, "/staysreviews/update/"+uuidStayReview.String(), strings.NewReader(""))

		router.HandleFunc("/staysreviews/update/{id}", hdl.UpdateStaysReview)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should be error updating stays review", func(t *testing.T) {
		r := httptest.NewRecorder()

		uuidStayReview, _ := uuid.NewV4()

		uuidStay, _ := uuid.NewV4()

		uuidUser, _ := uuid.NewV4()

		payload := staysreviews.StaysReviewEntity{
			StayID:      uuidStay,
			UserID:      uuidUser,
			Title:       "test",
			Description: "test",
			Rating:      1.2,
		}

		pBytes, _ := json.Marshal(payload)

		pBuf := bytes.NewBuffer(pBytes)

		svc.On("UpdateStaysReview", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("failed")).Once()

		req := httptest.NewRequest(http.MethodPut, "/staysreviews/update/"+uuidStayReview.String(), pBuf)

		router.HandleFunc("/staysreviews/update/{id}", hdl.UpdateStaysReview)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestStaysReviewsHandler_DeleteStayReviewHandler(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.StaysReviewsService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	t.Run("should be no errors", func(t *testing.T) {
		r := httptest.NewRecorder()

		uuidStayReview, _ := uuid.NewV4()

		svc.On("DeleteStaysReview", mock.Anything, mock.Anything).Return(nil).Once()

		req := httptest.NewRequest(http.MethodDelete, "/staysreviews/delete/"+uuidStayReview.String(), nil)

		router.HandleFunc("/staysreviews/delete/{id}", hdl.DeleteStaysReview)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})

	t.Run("should be error parsing uuid", func(t *testing.T) {
		r := httptest.NewRecorder()

		fakeID := "fake"

		req := httptest.NewRequest(http.MethodDelete, "/staysreviews/delete/"+fakeID, nil)

		router.HandleFunc("/staysreviews/delete/{id}", hdl.DeleteStaysReview)
		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should be error deleting stays review", func(t *testing.T) {
		r := httptest.NewRecorder()

		uuidStayReview, _ := uuid.NewV4()

		svc.On("DeleteStaysReview", mock.Anything, mock.Anything).Return(errors.New("failed")).Once()

		req := httptest.NewRequest(http.MethodDelete, "/staysreviews/delete/"+uuidStayReview.String(), nil)

		router.HandleFunc("/staysreviews/delete/{id}", hdl.DeleteStaysReview)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestStaysReviewsHandler_FindOneStayReviewHandler(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.StaysReviewsService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	t.Run("should be no errors", func(t *testing.T) {
		r := httptest.NewRecorder()

		uuID, _ := uuid.NewV4()

		expected := staysreviews.StaysReview{
			ID:          uuID,
			StayID:      uuID,
			UserID:      uuID,
			Title:       "test",
			Description: "test",
			Rating:      1.2,
		}

		svc.On("FindOneStaysReview", mock.Anything, mock.Anything).Return(&expected, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/staysreviews/"+uuID.String(), nil)

		router.HandleFunc("/staysreviews/{id}", hdl.FindOneStaysReview)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)

		assert.ObjectsAreEqual(&expected, r.Body)
	})

	t.Run("should be error parsing uuid", func(t *testing.T) {
		r := httptest.NewRecorder()
		fakeID := "fake"

		req := httptest.NewRequest(http.MethodGet, "/staysreviews/"+fakeID, nil)

		router.HandleFunc("/staysreviews/{id}", hdl.FindOneStaysReview)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should be error finding stays review", func(t *testing.T) {
		r := httptest.NewRecorder()

		uuID, _ := uuid.NewV4()

		svc.On("FindOneStaysReview", mock.Anything, mock.Anything).Return(nil, errors.New("failed")).Once()

		req := httptest.NewRequest(http.MethodGet, "/staysreviews/"+uuID.String(), nil)

		router.HandleFunc("/staysreviews/{id}", hdl.FindOneStaysReview)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestStaysReviewsHandler_FindAllStayReviews(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.StaysReviewsService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	t.Run("should be no errors", func(t *testing.T) {
		r := httptest.NewRecorder()

		uuID, _ := uuid.NewV4()

		expected := []staysreviews.StaysReview{
			{
				ID:          uuID,
				StayID:      uuID,
				UserID:      uuID,
				Title:       "test",
				Description: "test",
				Rating:      1.2,
			},
		}

		svc.On("FindAllStaysReviews", mock.Anything).Return(expected, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/staysreviews", nil)

		router.HandleFunc("/staysreviews", hdl.FindAllStaysReviews)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})

	t.Run("should be getting stays reviews", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("FindAllStaysReviews", mock.Anything).Return(nil, errors.New("failed")).Once()

		req := httptest.NewRequest(http.MethodGet, "/staysreviews", nil)

		router.HandleFunc("/staysreviews", hdl.FindAllStaysReviews)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}
