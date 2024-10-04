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
