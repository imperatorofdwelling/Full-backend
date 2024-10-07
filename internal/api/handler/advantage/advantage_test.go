package advantage

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/api/handler"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces/mocks"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/advantage"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"testing"
	"time"
)

func TestAdvantageHandler_NewAdvantageHandler(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.AdvantageService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	t.Run("should be no errors", func(t *testing.T) {
		hdl.NewAdvantageHandler(router)
	})
}

func TestAdvantageHandler_CreateAdvantage(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.AdvantageService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	testCases := []struct {
		title                 string
		payload               advantage.AdvantageEntity
		contentType           string
		contentTypeImg        string
		contentDispositionImg string
		statusCode            int
		error                 error
	}{
		{
			title: "should successfully create advantage",
			payload: advantage.AdvantageEntity{
				Title: "test",
				Image: createMockSvg(MaxAdvantageImgSize - 1000),
			},
			contentTypeImg:        "image/svg+xml",
			contentDispositionImg: `form-data; name="image"; filename="test.svg"`,
			statusCode:            http.StatusCreated,
			error:                 nil,
		},
		{
			title: "should be invalid image type",
			payload: advantage.AdvantageEntity{
				Title: "test",
				Image: createMockSvg(MaxAdvantageImgSize - 1000),
			},
			contentTypeImg:        "image/png",
			contentDispositionImg: `form-data; name="image"; filename="test.png"`,
			statusCode:            http.StatusBadRequest,
			error:                 handler.ErrImageTypeNotSvg,
		},
		{
			title: "should be invalid size",
			payload: advantage.AdvantageEntity{
				Title: "test",
				Image: createMockSvg(MaxAdvantageImgSize + 1000),
			},
			contentTypeImg:        "image/svg+xml",
			contentDispositionImg: `form-data; name="image"; filename="test.svg"`,
			statusCode:            http.StatusBadRequest,
			error:                 handler.ErrInvalidImageSize,
		},
		{
			title: "should be invalid Header Content-Type",
			payload: advantage.AdvantageEntity{
				Title: "test",
			},
			contentType: "application/json",
			statusCode:  http.StatusBadRequest,
			error:       nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			r := httptest.NewRecorder()

			var buf bytes.Buffer
			writer := multipart.NewWriter(&buf)

			err := writer.WriteField("title", tc.payload.Title)
			if err != nil {
				t.Fatal(err)
			}

			partHeader := make(textproto.MIMEHeader)
			partHeader.Set("Content-Type", tc.contentTypeImg)
			partHeader.Set("Content-Disposition", tc.contentDispositionImg)

			part, err := writer.CreatePart(partHeader)
			if err != nil {
				t.Fatal(err)
			}

			_, err = part.Write(tc.payload.Image)
			if err != nil {
				t.Fatal(err)
			}

			err = writer.Close()

			svc.On("CreateAdvantage", mock.Anything, mock.Anything).Return(nil, nil)

			req := httptest.NewRequest(http.MethodPost, "/advantage/create", &buf)

			if tc.contentType != "" {
				req.Header.Set("Content-Type", tc.contentType)
			} else {
				req.Header.Set("Content-Type", writer.FormDataContentType())
			}

			router.HandleFunc("/advantage/create", hdl.CreateAdvantage)

			router.ServeHTTP(r, req)

			assert.Equal(t, tc.statusCode, r.Code)

			if tc.error != nil {
				assert.Contains(t, r.Body.String(), tc.error.Error())
			}
		})
	}

	t.Run("should return error parsing form-data error", func(t *testing.T) {
		r := httptest.NewRecorder()

		var buf bytes.Buffer
		writer := multipart.NewWriter(&buf)

		for i := 0; i < 5; i++ {
			if err := writer.WriteField("field", "a"); err != nil {
				t.Fatalf("failed to write field: %v", err)
			}
		}

		err := writer.Close()
		assert.NoError(t, err)

		svc.On("CreateAdvantage", mock.Anything, mock.Anything).Return(nil, nil)

		req := httptest.NewRequest(http.MethodPost, "/advantage/create", &buf)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		router.HandleFunc("/advantage/create", hdl.CreateAdvantage)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
}

func TestAdvantageHandler_RemoveAdvantage(t *testing.T) {

	t.Run("should successfully remove advantage", func(t *testing.T) {
		log := logger.New(logger.EnvLocal)
		svc := mocks.AdvantageService{}
		hdl := Handler{
			Svc: &svc,
			Log: log,
		}
		router := chi.NewRouter()

		r := httptest.NewRecorder()

		advID, err := uuid.NewV4()
		if err != nil {
			t.Fatal(err)
		}

		advIDStr := advID.String()

		svc.On("RemoveAdvantage", context.Background(), advID).Return(nil)

		req := httptest.NewRequest(http.MethodDelete, "/advantage/"+advIDStr, nil)

		router.HandleFunc("/advantage/{advantageId}", hdl.RemoveAdvantage)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusNoContent, r.Code)

		svc.AssertExpectations(t)
	})

	t.Run("should be an error removing advantage", func(t *testing.T) {
		log := logger.New(logger.EnvLocal)
		svc := mocks.AdvantageService{}
		hdl := Handler{
			Svc: &svc,
			Log: log,
		}
		router := chi.NewRouter()

		r := httptest.NewRecorder()

		advID, err := uuid.NewV4()
		if err != nil {
			t.Fatal(err)
		}

		advIDStr := advID.String()

		svc.On("RemoveAdvantage", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(errors.New("failed to remove advantage"))

		req := httptest.NewRequest(http.MethodDelete, "/advantage/"+advIDStr, nil)

		router.HandleFunc("/advantage/{advantageId}", hdl.RemoveAdvantage)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)

		svc.AssertExpectations(t)
	})

	t.Run("should be fail when parsing uuid", func(t *testing.T) {
		log := logger.New(logger.EnvLocal)
		svc := mocks.AdvantageService{}
		hdl := Handler{
			Svc: &svc,
			Log: log,
		}
		router := chi.NewRouter()

		r := httptest.NewRecorder()

		invalidID := "invalid"

		svc.On("RemoveAdvantage", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(errors.New("failed to remove advantage"))

		req := httptest.NewRequest(http.MethodDelete, "/advantage/"+invalidID, nil)

		router.HandleFunc("/advantage/{advantageId}", hdl.RemoveAdvantage)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
}

func TestHandler_GetAllAdvantages(t *testing.T) {

	t.Run("should successfully get all advantages", func(t *testing.T) {
		log := logger.New(logger.EnvLocal)
		svc := mocks.AdvantageService{}
		hdl := Handler{
			Svc: &svc,
			Log: log,
		}
		router := chi.NewRouter()

		r := httptest.NewRecorder()

		id, err := uuid.NewV4()
		if err != nil {
			t.Fatal(err)
		}

		payload := []advantage.Advantage{
			{
				ID:        id,
				Title:     "test",
				Image:     "./assets/images/advantages/test.svg",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		svc.On("GetAllAdvantages", mock.Anything).Return(payload, nil)

		req := httptest.NewRequest(http.MethodGet, "/advantage/all", nil)

		router.HandleFunc("/advantage/all", hdl.GetAllAdvantages)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)

		var resp []advantage.Advantage

		err = json.Unmarshal(r.Body.Bytes(), &resp)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, payload[0].Title, resp[0].Title)
		assert.ObjectsAreEqualValues(payload, resp)

		svc.AssertExpectations(t)
	})

	t.Run("should fail while getting all advantages svc", func(t *testing.T) {
		log := logger.New(logger.EnvLocal)
		svc := mocks.AdvantageService{}
		hdl := Handler{
			Svc: &svc,
			Log: log,
		}
		router := chi.NewRouter()

		r := httptest.NewRecorder()

		svc.On("GetAllAdvantages", mock.Anything).Return(nil, errors.New("failed to get all advantages"))

		req := httptest.NewRequest(http.MethodGet, "/advantage/all", nil)

		router.HandleFunc("/advantage/all", hdl.GetAllAdvantages)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
		assert.Contains(t, r.Body.String(), "failed to get all advantages")

		svc.AssertExpectations(t)
	})
}

func TestHandler_UpdateAdvantage(t *testing.T) {
	id, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		title                 string
		id                    string
		payload               advantage.AdvantageEntity
		contentType           string
		contentTypeImg        string
		contentDispositionImg string
		statusCode            int
		error                 error
	}{
		{
			title: "should successfully update advantage",
			id:    id.String(),
			payload: advantage.AdvantageEntity{
				Title: "test",
				Image: createMockSvg(MaxAdvantageImgSize - 1000),
			},
			contentTypeImg:        "image/svg+xml",
			contentDispositionImg: `form-data; name="image"; filename="test.svg"`,
			statusCode:            http.StatusOK,
			error:                 nil,
		},
		{
			title: "should be invalid image type",
			id:    id.String(),
			payload: advantage.AdvantageEntity{
				Title: "test",
				Image: createMockSvg(MaxAdvantageImgSize - 1000),
			},
			contentTypeImg:        "image/png",
			contentDispositionImg: `form-data; name="image"; filename="test.png"`,
			statusCode:            http.StatusBadRequest,
			error:                 handler.ErrImageTypeNotSvg,
		},
		{
			title: "should be invalid size",
			id:    id.String(),
			payload: advantage.AdvantageEntity{
				Title: "test",
				Image: createMockSvg(MaxAdvantageImgSize + 1000),
			},
			contentTypeImg:        "image/svg+xml",
			contentDispositionImg: `form-data; name="image"; filename="test.svg"`,
			statusCode:            http.StatusBadRequest,
			error:                 handler.ErrInvalidImageSize,
		},
		{
			title: "should be invalid Header Content-Type",
			id:    id.String(),
			payload: advantage.AdvantageEntity{
				Title: "test",
			},
			contentType: "application/json",
			statusCode:  http.StatusBadRequest,
			error:       nil,
		},
		{
			title: "should be invalid id",
			id:    "invalid",
			payload: advantage.AdvantageEntity{
				Title: "test",
				Image: createMockSvg(MaxAdvantageImgSize - 1000),
			},
			contentTypeImg:        "image/svg+xml",
			contentDispositionImg: `form-data; name="image"; filename="test.svg"`,
			statusCode:            http.StatusBadRequest,
			error:                 nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.title, func(t *testing.T) {
			log := logger.New(logger.EnvLocal)
			svc := mocks.AdvantageService{}
			hdl := Handler{
				Svc: &svc,
				Log: log,
			}
			router := chi.NewRouter()
			r := httptest.NewRecorder()

			mockTime := time.Now()

			expected := advantage.Advantage{
				ID:        uuid.FromStringOrNil(testCase.id),
				Title:     testCase.payload.Title,
				Image:     "./assets/images/advantages/test.svg",
				CreatedAt: mockTime,
				UpdatedAt: mockTime,
			}

			svc.On("UpdateAdvantageByID", mock.Anything, mock.Anything, mock.Anything).Return(expected, nil)

			var buf bytes.Buffer
			writer := multipart.NewWriter(&buf)

			err = writer.WriteField("title", testCase.payload.Title)
			if err != nil {
				t.Fatal(err)
			}

			partHeader := make(textproto.MIMEHeader)
			partHeader.Set("Content-Type", testCase.contentTypeImg)
			partHeader.Set("Content-Disposition", testCase.contentDispositionImg)

			part, err := writer.CreatePart(partHeader)
			if err != nil {
				t.Fatal(err)
			}

			_, err = part.Write(testCase.payload.Image)
			if err != nil {
				t.Fatal(err)
			}

			err = writer.Close()

			req := httptest.NewRequest(http.MethodPatch, "/advantage/"+testCase.id, &buf)
			if testCase.contentType != "" {
				req.Header.Set("Content-Type", testCase.contentType)
			} else {
				req.Header.Set("Content-Type", writer.FormDataContentType())
			}

			router.HandleFunc("/advantage/{advantageId}", hdl.UpdateAdvantage)

			router.ServeHTTP(r, req)

			assert.Equal(t, testCase.statusCode, r.Code)

		})
	}
}

func createMockSvg(size int) []byte {
	//svg size is 358 bytes
	svg := `<?xml version="1.0" encoding="UTF-8"?>
		<svg version="1.1" id="svg1" width="100" height="100" xmlns="http://www.w3.org/2000/svg">
			<rect width="100" height="100" style="fill:blue" />
			<circle cx="50" cy="50" r="40" style="fill:yellow" />
			<text x="50" y="55" font-family="Verdana" font-size="20" fill="black" text-anchor="middle">Hello</text>
		</svg>`

	svgSize := size - len(svg)

	svgBytes := make([]byte, svgSize)

	svgBytes = append(svgBytes, []byte(svg)...)

	return svgBytes
}
