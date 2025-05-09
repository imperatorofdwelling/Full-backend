package stays

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/config"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces/mocks"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/stays"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/stays/amenity"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger"
	"github.com/imperatorofdwelling/Full-backend/pkg/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"image/png"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestStaysHandler_NewStaysHandler(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.StaysService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	t.Run("should be no errors private router", func(t *testing.T) {
		hdl.NewStaysHandler(router)
	})
}

func TestStaysHandler_CreateStay(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.StaysService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	fakeUUID, _ := uuid.NewV4()

	amenitiesMap := map[amenity.Amenity]bool{
		"Wi-fi":           true,
		"Air conditioner": false,
	}

	payload := stays.StayEntity{
		UserID:             fakeUUID,
		LocationID:         fakeUUID,
		Name:               "Luxury Apartment",
		Type:               stays.Apartment,
		Guests:             4,
		Amenities:          amenitiesMap,
		House:              "12B",
		Entrance:           "South",
		Address:            "123 Main Street",
		RoomsCount:         "3",
		BedsCount:          "2",
		Price:              "150",
		Period:             "day",
		OwnersRules:        "No smoking, no parties",
		CancellationPolicy: "Flexible",
		DescribeProperty:   "A beautiful apartment in downtown",
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
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.StaysService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	fakeUUID, _ := uuid.NewV4()

	amenitiesMap := map[amenity.Amenity]bool{
		"Wi-fi":           true,
		"Air conditioner": false,
	}

	expected := stays.Stay{
		ID:                 fakeUUID,
		UserID:             fakeUUID,
		LocationID:         fakeUUID,
		Name:               "Luxury Apartment",
		Type:               stays.Apartment,
		Guests:             4,
		Rating:             4.7,
		Amenities:          amenitiesMap,
		House:              "12B",
		Entrance:           "South",
		Address:            "123 Main Street",
		RoomsCount:         "3",
		BedsCount:          "2",
		Price:              "150",
		Period:             "day",
		OwnersRules:        "No smoking, no parties",
		CancellationPolicy: "Flexible",
		DescribeProperty:   "A beautiful apartment in downtown",
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
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

		assert.Equal(t, expected.UserID, expected.UserID)
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
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.StaysService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	fakeUUID, _ := uuid.NewV4()

	amenitiesMap := map[amenity.Amenity]bool{
		"Wi-fi":           true,
		"Air conditioner": false,
	}

	// Change from []*stays.Stay to []stays.StayResponse to match what the handler expects
	expected := []stays.StayResponse{
		{
			Stay: stays.Stay{
				ID:                 fakeUUID,
				UserID:             fakeUUID,
				LocationID:         fakeUUID,
				Name:               "Luxury Apartment",
				Type:               stays.Apartment,
				Guests:             4,
				Rating:             4.7,
				Amenities:          amenitiesMap,
				House:              "12B",
				Entrance:           "South",
				Address:            "123 Main Street",
				RoomsCount:         "3",
				BedsCount:          "2",
				Price:              "150",
				Period:             "day",
				OwnersRules:        "No smoking, no parties",
				CancellationPolicy: "Flexible",
				DescribeProperty:   "A beautiful apartment in downtown",
				CreatedAt:          time.Now(),
				UpdatedAt:          time.Now(),
			},
			Images: []stays.StayImage{}, // Empty slice for images if none are needed
		},
	}

	router.Get("/stays", hdl.GetStays)

	t.Run("should be no errors", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("GetStays", mock.Anything).Return(expected, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/stays", nil)

		router.ServeHTTP(r, req)

		var actual struct {
			Data []stays.StayResponse `json:"data"`
		}

		err := render.DecodeJSON(r.Body, &actual)
		if err != nil {
			t.Fatalf("Failed to decode JSON response: %v", err)
		}

		assert.Equal(t, http.StatusOK, r.Code)
		assert.Equal(t, expected[0].ID, actual.Data[0].ID)
	})

	t.Run("should be error getting stays", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("GetStays", mock.Anything).Return(nil, errors.New("failed to get stays")).Once()

		req := httptest.NewRequest(http.MethodGet, "/stays", nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestStaysHandler_DeleteStayByID(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
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
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.StaysService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	fakeUUID, _ := uuid.NewV4()

	amenitiesMap := map[amenity.Amenity]bool{
		"Wi-fi":           true,
		"Air conditioner": false,
	}

	payload := stays.StayEntity{
		UserID:             fakeUUID,
		LocationID:         fakeUUID,
		Name:               "Updated Luxury Apartment",
		Type:               stays.Apartment,
		Guests:             4,
		Amenities:          amenitiesMap,
		House:              "12B",
		Entrance:           "South",
		Address:            "123 Main Street",
		RoomsCount:         "3",
		BedsCount:          "2",
		Price:              "170",
		Period:             "day",
		OwnersRules:        "No smoking, no parties",
		CancellationPolicy: "Flexible",
		DescribeProperty:   "An updated beautiful apartment in downtown",
	}

	expected := stays.Stay{
		ID:                 fakeUUID,
		UserID:             fakeUUID,
		LocationID:         fakeUUID,
		Name:               "Updated Luxury Apartment",
		Type:               stays.Apartment,
		Guests:             4,
		Rating:             4.7,
		Amenities:          amenitiesMap,
		House:              "12B",
		Entrance:           "South",
		Address:            "123 Main Street",
		RoomsCount:         "3",
		BedsCount:          "2",
		Price:              "170",
		Period:             "day",
		OwnersRules:        "No smoking, no parties",
		CancellationPolicy: "Flexible",
		DescribeProperty:   "An updated beautiful apartment in downtown",
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
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
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.StaysService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	fakeUUID, _ := uuid.NewV4()

	amenitiesMap := map[amenity.Amenity]bool{
		"Wi-fi":           true,
		"Air conditioner": false,
	}

	expected := []*stays.Stay{
		{
			ID:                 fakeUUID,
			UserID:             fakeUUID,
			LocationID:         fakeUUID,
			Name:               "Luxury Apartment",
			Type:               stays.Apartment,
			Guests:             4,
			Rating:             4.7,
			Amenities:          amenitiesMap,
			House:              "12B",
			Entrance:           "South",
			Address:            "123 Main Street",
			RoomsCount:         "3",
			BedsCount:          "2",
			Price:              "150",
			Period:             "day",
			OwnersRules:        "No smoking, no parties",
			CancellationPolicy: "Flexible",
			DescribeProperty:   "A beautiful apartment in downtown",
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
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

func TestStaysHandler_GetStayImagesByStayID(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.StaysService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	fakeUUID, _ := uuid.NewV4()

	expected := []stays.StayImage{
		{
			ID:        fakeUUID,
			StayID:    fakeUUID,
			ImageName: "fakePath",
			IsMain:    false,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	t.Run("should be no errors", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("GetImagesByStayID", mock.Anything, mock.Anything).Return(expected, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/stays/images/"+fakeUUID.String(), nil)

		router.HandleFunc("/stays/images/{stayId}", hdl.GetStayImagesByStayID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)

		var actual struct {
			Data []stays.StayImage `json:"data"`
		}

		err := json.Unmarshal(r.Body.Bytes(), &actual)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, expected[0].ImageName, actual.Data[0].ImageName)
	})

	t.Run("should be parsing uuid error", func(t *testing.T) {
		r := httptest.NewRecorder()

		invalidUUID := "invalid"

		req := httptest.NewRequest(http.MethodGet, "/stays/images/"+invalidUUID, nil)

		router.HandleFunc("/stays/images/{imageId}", hdl.GetStayImagesByStayID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should be error getting stay images by stay id", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("GetImagesByStayID", mock.Anything, mock.Anything).Return(nil, errors.New("failed")).Once()

		req := httptest.NewRequest(http.MethodGet, "/stays/images/"+fakeUUID.String(), nil)

		router.HandleFunc("/stays/images/{stayId}", hdl.GetStayImagesByStayID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestStaysHandler_GetMainImageByStayID(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.StaysService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	fakeUUID, _ := uuid.NewV4()

	expected := stays.StayImage{
		ID:        fakeUUID,
		StayID:    fakeUUID,
		ImageName: "fakePath",
		IsMain:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("should be no errors", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("GetMainImageByStayID", mock.Anything, mock.Anything).Return(expected, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/stays/images/main/"+fakeUUID.String(), nil)

		router.HandleFunc("/stays/images/main/{stayId}", hdl.GetMainImageByStayID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)

		var actual struct {
			Data stays.StayImage `json:"data"`
		}

		err := json.Unmarshal(r.Body.Bytes(), &actual)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, expected.ImageName, actual.Data.ImageName)
	})

	t.Run("should be parsing uuid error", func(t *testing.T) {
		r := httptest.NewRecorder()

		invalidUUID := "invalid"

		req := httptest.NewRequest(http.MethodGet, "/stays/images/main/"+invalidUUID, nil)

		router.HandleFunc("/stays/images/main/{imageId}", hdl.GetMainImageByStayID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should be error getting stay images by stay id", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("GetMainImageByStayID", mock.Anything, mock.Anything).Return(stays.StayImage{}, errors.New("failed")).Once()

		req := httptest.NewRequest(http.MethodGet, "/stays/images/main/"+fakeUUID.String(), nil)

		router.HandleFunc("/stays/images/main/{stayId}", hdl.GetMainImageByStayID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestStaysHandler_CreateImages(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.StaysService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	mockUUID, _ := uuid.NewV4()

	t.Run("should be no errors", func(t *testing.T) {
		r := httptest.NewRecorder()

		var buf bytes.Buffer
		writer := multipart.NewWriter(&buf)

		_ = writer.WriteField("stay_id", mockUUID.String())

		mockImg := testhelper.CreateMockPng()

		part, err := writer.CreateFormFile("images", "test")
		if err != nil {
			t.Fatal(err)
		}

		err = png.Encode(part, mockImg)
		if err != nil {
			t.Fatal(err)
		}

		err = writer.Close()
		if err != nil {
			t.Fatal(err)
		}

		svc.On("CreateImages", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

		req := httptest.NewRequest(http.MethodPost, "/stays/images", &buf)
		req.Header.Add("Content-Type", writer.FormDataContentType())

		router.HandleFunc("/stays/images", hdl.CreateImages)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusCreated, r.Code)
	})

	t.Run("should be error creating images", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("CreateImages", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("failed")).Once()

		var buf bytes.Buffer
		writer := multipart.NewWriter(&buf)

		_ = writer.WriteField("stay_id", mockUUID.String())

		mockImg := testhelper.CreateMockPng()

		part, err := writer.CreateFormFile("images", "test")
		if err != nil {
			t.Fatal(err)
		}

		err = png.Encode(part, mockImg)
		if err != nil {
			t.Fatal(err)
		}

		err = writer.Close()
		if err != nil {
			t.Fatal(err)
		}

		req := httptest.NewRequest(http.MethodPost, "/stays/images", &buf)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		router.HandleFunc("/stays/images", hdl.CreateImages)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})

	t.Run("should be parsing uuid error", func(t *testing.T) {
		r := httptest.NewRecorder()

		invalidUUID := "invalid"

		var buf bytes.Buffer
		writer := multipart.NewWriter(&buf)

		_ = writer.WriteField("stay_id", invalidUUID)

		mockImg := testhelper.CreateMockPng()

		part, err := writer.CreateFormFile("images", "test")
		if err != nil {
			t.Fatal(err)
		}

		err = png.Encode(part, mockImg)
		if err != nil {
			t.Fatal(err)
		}

		err = writer.Close()
		if err != nil {
			t.Fatal(err)
		}

		svc.On("CreateImages", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

		req := httptest.NewRequest(http.MethodPost, "/stays/images", &buf)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		router.HandleFunc("/stays/images", hdl.CreateImages)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should be error parsing multipart form data", func(t *testing.T) {
		r := httptest.NewRecorder()

		var buf bytes.Buffer

		svc.On("CreateImages", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

		req := httptest.NewRequest(http.MethodPost, "/stays/images", &buf)

		router.HandleFunc("/stays/images", hdl.CreateImages)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
}

func TestStaysHandler_DeleteStayImage(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.StaysService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	fakeUUID, _ := uuid.NewV4()

	t.Run("should be no errors", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("DeleteStayImage", mock.Anything, mock.Anything).Return(nil).Once()

		req := httptest.NewRequest(http.MethodDelete, "/stays/images/delete/"+fakeUUID.String(), nil)

		router.HandleFunc("/stays/images/delete/{imageId}", hdl.DeleteStayImage)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusNoContent, r.Code)
	})

	t.Run("should be parsing uuid error", func(t *testing.T) {
		r := httptest.NewRecorder()

		invalidUUID := "invalid"

		req := httptest.NewRequest(http.MethodDelete, "/stays/images/delete/"+invalidUUID, nil)

		router.HandleFunc("/stays/images/delete/{imageId}", hdl.DeleteStayImage)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should be error deleting stay image", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("DeleteStayImage", mock.Anything, mock.Anything).Return(errors.New("failed")).Once()

		req := httptest.NewRequest(http.MethodDelete, "/stays/images/delete/"+fakeUUID.String(), nil)

		router.HandleFunc("/stays/images/delete/{imageId}", hdl.DeleteStayImage)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestStaysHandler_CreateMainImage(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.StaysService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	mockUUID, _ := uuid.NewV4()

	t.Run("should be no errors", func(t *testing.T) {
		r := httptest.NewRecorder()

		var buf bytes.Buffer
		writer := multipart.NewWriter(&buf)

		_ = writer.WriteField("stay_id", mockUUID.String())

		mockImg := testhelper.CreateMockPng()

		part, err := writer.CreateFormFile("images", "test")
		if err != nil {
			t.Fatal(err)
		}

		err = png.Encode(part, mockImg)
		if err != nil {
			t.Fatal(err)
		}

		err = writer.Close()
		if err != nil {
			t.Fatal(err)
		}

		svc.On("CreateMainImage", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

		req := httptest.NewRequest(http.MethodPost, "/stays/images/main", &buf)
		req.Header.Add("Content-Type", writer.FormDataContentType())

		router.HandleFunc("/stays/images/main", hdl.CreateMainImage)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusCreated, r.Code)
	})

	t.Run("should be parsing uuid error", func(t *testing.T) {
		r := httptest.NewRecorder()

		invalidUUID := "invalid"

		var buf bytes.Buffer
		writer := multipart.NewWriter(&buf)

		_ = writer.WriteField("stay_id", invalidUUID)

		mockImg := testhelper.CreateMockPng()

		part, err := writer.CreateFormFile("images", "test")
		if err != nil {
			t.Fatal(err)
		}

		err = png.Encode(part, mockImg)
		if err != nil {
			t.Fatal(err)
		}

		err = writer.Close()
		if err != nil {
			t.Fatal(err)
		}

		req := httptest.NewRequest(http.MethodPost, "/stays/images/main", &buf)
		req.Header.Add("Content-Type", writer.FormDataContentType())

		router.HandleFunc("/stays/images/main", hdl.CreateMainImage)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should be error creating main image", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("CreateMainImage", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("failed")).Once()

		var buf bytes.Buffer
		writer := multipart.NewWriter(&buf)

		_ = writer.WriteField("stay_id", mockUUID.String())

		mockImg := testhelper.CreateMockPng()

		part, err := writer.CreateFormFile("images", "test")
		if err != nil {
			t.Fatal(err)
		}

		err = png.Encode(part, mockImg)
		if err != nil {
			t.Fatal(err)
		}

		err = writer.Close()
		if err != nil {
			t.Fatal(err)
		}

		req := httptest.NewRequest(http.MethodPost, "/stays/images/main", &buf)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		router.HandleFunc("/stays/images/main", hdl.CreateMainImage)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})

	t.Run("should be error parsing multipart form data", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("CreateMainImage", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

		req := httptest.NewRequest(http.MethodPost, "/stays/images/main", bytes.NewReader([]byte{}))

		router.HandleFunc("/stays/images/main", hdl.CreateMainImage)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
}

func TestStaysHandler_GetStaysByLocationID(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.StaysService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	mockUUID, _ := uuid.NewV4()

	// Updated mockStay to match the new Stay structure
	mockStay := []stays.Stay{
		{
			ID:                 mockUUID,
			UserID:             mockUUID,
			LocationID:         mockUUID,
			Name:               "Luxurious Apartment in City Center",
			Type:               stays.Apartment,
			Guests:             6,
			Rating:             4.8,
			Amenities:          map[amenity.Amenity]bool{"Wi-fi": true, "Air conditioner": true},
			House:              "22A",
			Entrance:           "North Entrance",
			Address:            "Main Street",
			RoomsCount:         "2",
			BedsCount:          "3",
			Price:              "120.0",
			Period:             "day",
			OwnersRules:        "No smoking, no pets",
			CancellationPolicy: "Free cancellation 24 hours before check-in",
			DescribeProperty:   "Beautiful apartment in the heart of the city",
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		},
	}

	t.Run("should be no errors", func(t *testing.T) {
		r := httptest.NewRecorder()
		svc.On("GetStaysByLocationID", mock.Anything, mock.Anything).Return(&mockStay, nil).Once()
		req := httptest.NewRequest(http.MethodGet, "/stays/location/"+mockUUID.String(), nil)
		router.HandleFunc("/stays/location/{locationId}", hdl.GetStaysByLocationID)
		router.ServeHTTP(r, req)
		assert.Equal(t, http.StatusOK, r.Code)
	})

	t.Run("should be parsing uuid error", func(t *testing.T) {
		r := httptest.NewRecorder()
		invalidUUID := "invalid"
		req := httptest.NewRequest(http.MethodGet, "/stays/location/"+invalidUUID, nil)
		router.HandleFunc("/stays/location/{locationId}", hdl.GetStaysByLocationID)
		router.ServeHTTP(r, req)
		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should be error getting stays by locationID", func(t *testing.T) {
		r := httptest.NewRecorder()
		errMessage := "error message"
		svc.On("GetStaysByLocationID", mock.Anything, mock.Anything).Return(nil, errors.New(errMessage)).Once()
		req := httptest.NewRequest(http.MethodGet, "/stays/location/"+mockUUID.String(), nil)
		router.HandleFunc("/stays/location/{locationId}", hdl.GetStaysByLocationID)
		router.ServeHTTP(r, req)
		assert.Equal(t, http.StatusInternalServerError, r.Code)

	})
}
