package stays

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces/mocks"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/stays"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestStaysHandler_NewStaysHandler(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.StaysService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	t.Run("should be no errors", func(t *testing.T) {
		hdl.NewStaysHandler(router)
	})
}

func TestStaysHandler_CreateStay(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.StaysService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	fakeUUID, _ := uuid.NewV4()

	payload := stays.StayEntity{
		Entrance:            "string",
		Floor:               "string",
		Guests:              0,
		House:               "string",
		IsSmokingProhibited: false,
		LocationID:          fakeUUID,
		Name:                "string",
		NumberOfBathrooms:   0,
		NumberOfBedrooms:    0,
		NumberOfBeds:        0,
		Price:               0,
		Room:                "string",
		Square:              0,
		Street:              "string",
		Type:                "string",
		UserID:              fakeUUID,
	}

	pMarshalled, _ := json.Marshal(payload)

	t.Run("should be no errors", func(t *testing.T) {
		r := httptest.NewRecorder()

		pBuf := bytes.NewBuffer(pMarshalled)

		svc.On("CreateStay", mock.Anything, mock.Anything).Return(nil).Once()

		req := httptest.NewRequest(http.MethodPost, "/stays/create", pBuf)

		router.HandleFunc("/stays/create", hdl.CreateStay)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusCreated, r.Code)
	})

	t.Run("should be decode error", func(t *testing.T) {
		r := httptest.NewRecorder()

		invalidBody := strings.NewReader("")

		req := httptest.NewRequest(http.MethodPost, "/stays/create", invalidBody)

		router.HandleFunc("/stays/create", hdl.CreateStay)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should be error creating stay", func(t *testing.T) {
		r := httptest.NewRecorder()

		pBuf := bytes.NewBuffer(pMarshalled)

		svc.On("CreateStay", mock.Anything, mock.Anything).Return(errors.New("failed to create stay")).Once()

		req := httptest.NewRequest(http.MethodPost, "/stays/create", pBuf)

		router.HandleFunc("/stays/create", hdl.CreateStay)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestStaysHandler_GetStayByID(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.StaysService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	fakeUUID, _ := uuid.NewV4()

	expected := stays.Stay{
		Entrance:            "string",
		Floor:               "string",
		Guests:              0,
		House:               "string",
		IsSmokingProhibited: false,
		LocationID:          fakeUUID,
		Name:                "string",
		NumberOfBathrooms:   0,
		NumberOfBedrooms:    0,
		NumberOfBeds:        0,
		Price:               0,
		Room:                "string",
		Square:              0,
		Street:              "string",
		Type:                "string",
		UserID:              fakeUUID,
	}

	t.Run("should be no errors", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("GetStayByID", mock.Anything, mock.Anything).Return(&expected, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/stays/"+fakeUUID.String(), nil)

		router.HandleFunc("/stays/{stayId}", hdl.GetStayByID)

		router.ServeHTTP(r, req)

		var actual stays.Stay

		_ = render.DecodeJSON(r.Body, &actual)

		assert.Equal(t, http.StatusOK, r.Code)

		assert.Equal(t, expected.UserID, actual.UserID)
	})

	t.Run("should be uuid parsing error", func(t *testing.T) {
		r := httptest.NewRecorder()

		invalidUUID := "invalid"

		req := httptest.NewRequest(http.MethodGet, "/stays/"+invalidUUID, nil)

		router.HandleFunc("/stays/{stayId}", hdl.GetStayByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should be error getting stay", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("GetStayByID", mock.Anything, mock.Anything).Return(nil, errors.New("failed to get stay")).Once()

		req := httptest.NewRequest(http.MethodGet, "/stays/"+fakeUUID.String(), nil)

		router.HandleFunc("/stays/{stayId}", hdl.GetStayByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestStaysHandler_GetStays(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.StaysService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	fakeUUID, _ := uuid.NewV4()

	expected := []*stays.Stay{
		{
			ID:                  fakeUUID,
			Entrance:            "string",
			Floor:               "string",
			Guests:              0,
			House:               "string",
			ImageMain:           "string",
			Images:              []string{"ssdsd"},
			IsSmokingProhibited: false,
			LocationID:          fakeUUID,
			Name:                "string",
			NumberOfBathrooms:   0,
			NumberOfBedrooms:    0,
			NumberOfBeds:        0,
			Price:               2.5,
			Room:                "string",
			Square:              0,
			Street:              "string",
			Type:                stays.Apartment,
			UserID:              fakeUUID,
			CreatedAt:           time.Now(),
			UpdatedAt:           time.Now(),
			Rating:              1.1,
		},
	}

	t.Run("should be no errors", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("GetStays", mock.Anything).Return(expected, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/stays", nil)

		router.HandleFunc("/stays", hdl.GetStays)

		router.ServeHTTP(r, req)

		var actual []stays.Stay

		_ = render.DecodeJSON(r.Body, &actual)

		assert.Equal(t, http.StatusOK, r.Code)

		assert.Equal(t, expected[0].ID, actual[0].ID)
	})

	t.Run("should be error getting stays", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("GetStays", mock.Anything).Return(nil, errors.New("failed to get stays")).Once()

		req := httptest.NewRequest(http.MethodGet, "/stays", nil)

		router.HandleFunc("/stays", hdl.GetStays)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestStaysHandler_DeleteStayByID(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.StaysService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	fakeUUID, _ := uuid.NewV4()

	t.Run("should be no errors", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("DeleteStayByID", mock.Anything, mock.Anything).Return(nil).Once()

		req := httptest.NewRequest(http.MethodDelete, "/stays/"+fakeUUID.String(), nil)

		router.HandleFunc("/stays/{stayId}", hdl.DeleteStayByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusNoContent, r.Code)
	})

	t.Run("should be parsing uuid error", func(t *testing.T) {
		r := httptest.NewRecorder()

		invalidUUID := "invalid"

		req := httptest.NewRequest(http.MethodDelete, "/stays/"+invalidUUID, nil)

		router.HandleFunc("/stays/{stayId}", hdl.DeleteStayByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should be error deleting stay", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("DeleteStayByID", mock.Anything, mock.Anything).Return(errors.New("failed to delete stay")).Once()

		req := httptest.NewRequest(http.MethodDelete, "/stays/"+fakeUUID.String(), nil)

		router.HandleFunc("/stays/{stayId}", hdl.DeleteStayByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestStaysHandler_UpdateStayByID(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.StaysService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	fakeUUID, _ := uuid.NewV4()

	payload := stays.StayEntity{
		Entrance:            "string",
		Floor:               "string",
		Guests:              0,
		House:               "string",
		IsSmokingProhibited: false,
		LocationID:          fakeUUID,
		Name:                "string",
		NumberOfBathrooms:   0,
		NumberOfBedrooms:    0,
		NumberOfBeds:        0,
		Price:               0,
		Room:                "string",
		Square:              0,
		Street:              "string",
		Type:                "string",
		UserID:              fakeUUID,
	}

	expected := stays.Stay{
		ID:                  fakeUUID,
		Entrance:            "string",
		Floor:               "string",
		Guests:              0,
		House:               "string",
		IsSmokingProhibited: false,
		LocationID:          fakeUUID,
		Name:                "string",
		NumberOfBathrooms:   0,
		NumberOfBedrooms:    0,
		NumberOfBeds:        0,
		Price:               0,
		Room:                "string",
		Square:              0,
		Street:              "string",
		Type:                "string",
		UserID:              fakeUUID,
	}

	pMarshalled, _ := json.Marshal(payload)

	t.Run("should be no errors", func(t *testing.T) {
		r := httptest.NewRecorder()

		pBuf := bytes.NewBuffer(pMarshalled)

		svc.On("UpdateStayByID", mock.Anything, mock.Anything, mock.Anything).Return(&expected, nil).Once()

		req := httptest.NewRequest(http.MethodPost, "/stays/"+fakeUUID.String(), pBuf)

		router.HandleFunc("/stays/{stayId}", hdl.UpdateStayByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})

	t.Run("should be error parsing uuid", func(t *testing.T) {
		r := httptest.NewRecorder()

		invalidUUID := "invalid"

		req := httptest.NewRequest(http.MethodPost, "/stays/"+invalidUUID, nil)

		router.HandleFunc("/stays/{stayId}", hdl.UpdateStayByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should be error decoding json", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPost, "/stays/"+fakeUUID.String(), strings.NewReader(""))

		router.HandleFunc("/stays/{stayId}", hdl.UpdateStayByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should be error updating stay", func(t *testing.T) {
		r := httptest.NewRecorder()

		pBuf := bytes.NewBuffer(pMarshalled)

		svc.On("UpdateStayByID", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("failed")).Once()

		req := httptest.NewRequest(http.MethodPut, "/stays/"+fakeUUID.String(), pBuf)

		router.HandleFunc("/stays/{stayId}", hdl.UpdateStayByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestStaysHandler_GetStaysByUserID(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.StaysService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	fakeUUID, _ := uuid.NewV4()

	expected := []*stays.Stay{
		{
			ID:                  fakeUUID,
			Entrance:            "string",
			Floor:               "string",
			Guests:              0,
			House:               "string",
			ImageMain:           "string",
			Images:              []string{"ssdsd"},
			IsSmokingProhibited: false,
			LocationID:          fakeUUID,
			Name:                "string",
			NumberOfBathrooms:   0,
			NumberOfBedrooms:    0,
			NumberOfBeds:        0,
			Price:               2.5,
			Room:                "string",
			Square:              0,
			Street:              "string",
			Type:                stays.Apartment,
			UserID:              fakeUUID,
			CreatedAt:           time.Now(),
			UpdatedAt:           time.Now(),
			Rating:              1.1,
		},
	}

	t.Run("should be no errors", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("GetStaysByUserID", mock.Anything, mock.Anything).Return(expected, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/stays/user/"+fakeUUID.String(), nil)

		router.HandleFunc("/stays/user/{userId}", hdl.GetStaysByUserID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})

	t.Run("should be parsing uuid error", func(t *testing.T) {
		r := httptest.NewRecorder()

		invalidUUID := "invalid"

		req := httptest.NewRequest(http.MethodGet, "/stays/user/"+invalidUUID, nil)

		router.HandleFunc("/stays/user/{userId}", hdl.GetStaysByUserID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should be error getting stays by user id", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("GetStaysByUserID", mock.Anything, mock.Anything).Return(nil, errors.New("failed")).Once()

		req := httptest.NewRequest(http.MethodGet, "/stays/user/"+fakeUUID.String(), nil)

		router.HandleFunc("/stays/user/{userId}", hdl.GetStaysByUserID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})

}
