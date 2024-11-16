package location

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces/mocks"
	models "github.com/imperatorofdwelling/Full-backend/internal/domain/models/location"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLocationHandler_FindByNameMatch(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.LocationService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()
	hdl.NewLocationHandler(router)

	t.Run("should be correct response", func(t *testing.T) {
		r := httptest.NewRecorder()
		payload := "алейск"
		expected := &[]models.Location{
			{
				City:            "Алейск",
				FederalDistrict: "Сибирский",
				FiasID:          "ae716080-f27b-40b6-a555-cf8b518e849e",
				KladrID:         "2200000200000",
				Lat:             "52.4922513",
				Lon:             "82.7793606",
				Okato:           "1403000000",
				Oktmo:           "1703000001",
				Population:      29,
				RegionIsoCode:   "RU-ALT",
				RegionName:      "Алтайский край",
			},
		}

		svc.On("FindByNameMatch", mock.Anything, mock.Anything).Return(expected, nil).Once()

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/locations/match/%s", payload), nil)
		assert.NoError(t, err)

		router.HandleFunc("/locations/match/{locationName}", hdl.FindByNameMatch)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)

		var actual struct {
			Data []models.Location `json:"data"`
		}

		err = render.DecodeJSON(r.Body, &actual)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, expected, &actual.Data)
	})

	t.Run("should be error response", func(t *testing.T) {
		r := httptest.NewRecorder()
		payload := "invalid"

		expectedErr := fmt.Errorf("location not found")

		svc.On("FindByNameMatch", context.Background(), payload).Return(nil, expectedErr).Once()

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/locations/match/%s", payload), nil)
		assert.NoError(t, err)

		router := chi.NewRouter()

		router.HandleFunc("/locations/match/{locationName}", hdl.FindByNameMatch)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)

		var response responseApi.ResponseError

		err = json.Unmarshal(r.Body.Bytes(), &response)

		assert.NoError(t, err)
		assert.Contains(t, response.Error, expectedErr.Error())
		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}
