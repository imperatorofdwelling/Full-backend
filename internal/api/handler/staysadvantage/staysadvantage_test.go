package staysadvantage

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/config"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces/mocks"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/staysadvantage"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestStaysAdvantagesHandler_NewStaysAdvantageHandler(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.StaysAdvantageService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()
	t.Run("should be no error", func(t *testing.T) {
		hdl.NewStaysAdvantageHandler(router)
	})
}

func TestStaysAdvantagesHandler_CreateStaysAdvantage(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.StaysAdvantageService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	fakeUUID, _ := uuid.NewV4()

	payload := staysadvantage.StayAdvantageCreateReq{
		StayID:      fakeUUID,
		AdvantageID: fakeUUID,
	}

	pBytes, _ := json.Marshal(payload)

	t.Run("should be no error", func(t *testing.T) {
		r := httptest.NewRecorder()

		pBuf := bytes.NewBuffer(pBytes)

		svc.On("CreateStaysAdvantage", mock.Anything, mock.Anything).Return(nil).Once()

		req := httptest.NewRequest(http.MethodPost, "/staysadvantage/create", pBuf)

		router.HandleFunc("/staysadvantage/create", hdl.CreateStaysAdvantage)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusCreated, r.Code)
	})

	t.Run("should be error decoding body", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPost, "/staysadvantage/create", strings.NewReader(""))

		router.HandleFunc("/staysadvantage/create", hdl.CreateStaysAdvantage)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should be error creating stays advantage", func(t *testing.T) {
		r := httptest.NewRecorder()

		pBuf := bytes.NewBuffer(pBytes)

		svc.On("CreateStaysAdvantage", mock.Anything, mock.Anything).Return(errors.New("error")).Once()

		req := httptest.NewRequest(http.MethodPost, "/staysadvantage/create", pBuf)

		router.HandleFunc("/staysadvantage/create", hdl.CreateStaysAdvantage)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestStaysAdvantagesHandler_DeleteStaysAdvantageByID(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.StaysAdvantageService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	fakeUUID, _ := uuid.NewV4()

	t.Run("should be no error", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("DeleteStaysAdvantageByID", mock.Anything, mock.Anything).Return(nil).Once()

		req := httptest.NewRequest(http.MethodDelete, "/staysadvantage/"+fakeUUID.String(), nil)

		router.HandleFunc("/staysadvantage/{id}", hdl.DeleteStaysAdvantageByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusNoContent, r.Code)
	})

	t.Run("should be error parsing uuid", func(t *testing.T) {
		r := httptest.NewRecorder()

		invalidUUID := "invalid"

		req := httptest.NewRequest(http.MethodDelete, "/staysadvantage/"+invalidUUID, nil)

		router.HandleFunc("/staysadvantage/{id}", hdl.DeleteStaysAdvantageByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should be error deleting stays advantage", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("DeleteStaysAdvantageByID", mock.Anything, mock.Anything).Return(errors.New("error")).Once()

		req := httptest.NewRequest(http.MethodDelete, "/staysadvantage/"+fakeUUID.String(), nil)

		router.HandleFunc("/staysadvantage/{id}", hdl.DeleteStaysAdvantageByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}
