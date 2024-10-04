package reservation

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces/mocks"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/reservation"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestReservationHandler_NewReservationHandler(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.ReservationService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	t.Run("should be no errors", func(t *testing.T) {
		hdl.NewReservationHandler(router)
	})
}

func TestReservationHandler_CreateReservation(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.ReservationService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	fakeUUID, _ := uuid.NewV4()

	payload := reservation.ReservationEntity{
		StayID:    fakeUUID,
		UserID:    fakeUUID,
		Arrived:   time.Now(),
		Departure: time.Now(),
	}

	pBytes, _ := json.Marshal(payload)

	t.Run("should be no errors", func(t *testing.T) {
		r := httptest.NewRecorder()

		pBuf := bytes.NewBuffer(pBytes)

		svc.On("CreateReservation", mock.Anything, mock.Anything).Return(nil).Once()

		req := httptest.NewRequest(http.MethodPost, "/reservation/create", pBuf)

		router.HandleFunc("/reservation/create", hdl.CreateReservation)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusCreated, r.Code)
	})

	t.Run("should be error decoding body", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPost, "/reservation/create", strings.NewReader(""))

		router.HandleFunc("/reservation/create", hdl.CreateReservation)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should be error creating reservation", func(t *testing.T) {
		r := httptest.NewRecorder()

		pBuf := bytes.NewBuffer(pBytes)

		svc.On("CreateReservation", mock.Anything, mock.Anything).Return(errors.New("failed")).Once()

		req := httptest.NewRequest(http.MethodPost, "/reservation/create", pBuf)

		router.HandleFunc("/reservation/create", hdl.CreateReservation)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestReservationHandler_UpdateReservation(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.ReservationService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	fakeUUID, _ := uuid.NewV4()

	payload := reservation.ReservationEntity{
		StayID:    fakeUUID,
		UserID:    fakeUUID,
		Arrived:   time.Now(),
		Departure: time.Now(),
	}

	pBytes, _ := json.Marshal(payload)

	t.Run("should be no errors", func(t *testing.T) {
		r := httptest.NewRecorder()

		pBuf := bytes.NewBuffer(pBytes)

		svc.On("UpdateReservation", mock.Anything, mock.Anything).Return(nil).Once()

		req := httptest.NewRequest(http.MethodPut, "/reservation/"+fakeUUID.String(), pBuf)

		router.HandleFunc("/reservation/{reservationId}", hdl.UpdateReservation)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})

	t.Run("should be error decoding body", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPut, "/reservation/"+fakeUUID.String(), strings.NewReader(""))

		router.HandleFunc("/reservation/{reservationId}", hdl.UpdateReservation)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should be error updating reservation", func(t *testing.T) {
		r := httptest.NewRecorder()

		pBuf := bytes.NewBuffer(pBytes)

		svc.On("UpdateReservation", mock.Anything, mock.Anything).Return(errors.New("failed")).Once()

		req := httptest.NewRequest(http.MethodPut, "/reservation/"+fakeUUID.String(), pBuf)

		router.HandleFunc("/reservation/{reservationId}", hdl.UpdateReservation)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestReservationHandler_DeleteReservationByID(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.ReservationService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	fakeUUID, _ := uuid.NewV4()

	t.Run("should be no errors", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("DeleteReservationByID", mock.Anything, mock.Anything).Return(nil).Once()

		req := httptest.NewRequest(http.MethodDelete, "/reservation/"+fakeUUID.String(), nil)

		router.HandleFunc("/reservation/{reservationID}", hdl.DeleteReservationByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})

	t.Run("should be error parsing uuid", func(t *testing.T) {
		r := httptest.NewRecorder()

		invalidUUID := "invalid"

		req := httptest.NewRequest(http.MethodDelete, "/reservation/"+invalidUUID, nil)

		router.HandleFunc("/reservation/{reservationID}", hdl.DeleteReservationByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should be error deleting reservation", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("DeleteReservationByID", mock.Anything, mock.Anything).Return(errors.New("failed")).Once()

		req := httptest.NewRequest(http.MethodDelete, "/reservation/"+fakeUUID.String(), nil)

		router.HandleFunc("/reservation/{reservationID}", hdl.DeleteReservationByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestReservationHandler_GetReservationByID(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.ReservationService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	fakeUUID, _ := uuid.NewV4()

	expected := &reservation.Reservation{
		ID:        fakeUUID,
		StayID:    fakeUUID,
		UserID:    fakeUUID,
		Arrived:   time.Now(),
		Departure: time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("should be no errors", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("GetReservationByID", mock.Anything, mock.Anything).Return(expected, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/reservation/"+fakeUUID.String(), nil)

		router.HandleFunc("/reservation/{reservationID}", hdl.GetReservationByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})

	t.Run("should be error parsing uuid", func(t *testing.T) {
		r := httptest.NewRecorder()

		invalidUUID := "invalid"

		req := httptest.NewRequest(http.MethodGet, "/reservation/"+invalidUUID, nil)

		router.HandleFunc("/reservation/{reservationID}", hdl.GetReservationByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should be error getting reservation", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("GetReservationByID", mock.Anything, mock.Anything).Return(nil, errors.New("failed")).Once()

		req := httptest.NewRequest(http.MethodGet, "/reservation/"+fakeUUID.String(), nil)

		router.HandleFunc("/reservation/{reservationID}", hdl.GetReservationByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestReservationHandler_GetAllReservationsByUser(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.ReservationService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	fakeUUID, _ := uuid.NewV4()

	expected := []reservation.Reservation{
		{
			ID:        fakeUUID,
			StayID:    fakeUUID,
			UserID:    fakeUUID,
			Arrived:   time.Now(),
			Departure: time.Now(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	t.Run("should be no errors", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("GetAllReservationsByUser", mock.Anything, mock.Anything).Return(&expected, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/reservation/user/"+fakeUUID.String(), nil)

		router.HandleFunc("/reservation/user/{userID}", hdl.GetAllReservationsByUser)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})

	t.Run("should be error parsing uuid", func(t *testing.T) {
		r := httptest.NewRecorder()

		invalidUUID := "invalid"

		req := httptest.NewRequest(http.MethodGet, "/reservation/user/"+invalidUUID, nil)

		router.HandleFunc("/reservation/user/{userID}", hdl.GetAllReservationsByUser)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should be error getting reservation", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("GetAllReservationsByUser", mock.Anything, mock.Anything).Return(nil, errors.New("failed")).Once()

		req := httptest.NewRequest(http.MethodGet, "/reservation/user/"+fakeUUID.String(), nil)

		router.HandleFunc("/reservation/user/{userID}", hdl.GetAllReservationsByUser)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}
