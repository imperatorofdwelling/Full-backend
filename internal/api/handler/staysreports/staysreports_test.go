package staysreports

import (
	"bytes"
	"database/sql"
	"errors"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	handler "github.com/imperatorofdwelling/Full-backend/internal/api/handler/user"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces/mocks"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"strings"
	"testing"
	"time"
)

func TestStaysReportsHandler_NewStaysReportsHandler(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.StaysReportsService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	t.Run("should be no errors", func(t *testing.T) {
		hdl.NewStaysReportsHandler(router)
	})
}

func TestStaysReportsHandler_UserIdError(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.StaysReportsService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	router.Get("/report", hdl.GetAllStaysReports)
	router.Get("/report/{stayId}", hdl.GetStaysReportById)
	router.Post("/report/create", hdl.CreateStaysReports)
	router.Patch("/report/{reportId}", hdl.UpdateStaysReports)
	router.Delete("/report/{reportId}", hdl.DeleteStaysReports)

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("get all user error", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/report", nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)
		assert.Equal(t, http.StatusUnauthorized, r.Code)
	})
	t.Run("get one user error", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/report/"+testUserID.String(), nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)
		assert.Equal(t, http.StatusUnauthorized, r.Code)
	})
	t.Run("post user error", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPost, "/report/create", nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)
		assert.Equal(t, http.StatusUnauthorized, r.Code)
	})
	t.Run("delete user error", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodDelete, "/report/"+testUserID.String(), nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)
		assert.Equal(t, http.StatusUnauthorized, r.Code)
	})
	t.Run("patch user error", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPatch, "/report/"+testUserID.String(), nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)
		assert.Equal(t, http.StatusUnauthorized, r.Code)
	})
}

func TestStaysReportsHandler_ParamsError(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.StaysReportsService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	router.Use(handler.JWTMiddleware("your-secret-key", log))

	router.Post("/report/create/{stayId}", hdl.CreateStaysReports)
	router.Delete("/report/{reportId}", hdl.DeleteStaysReports)
	router.Patch("/report/{reportId}", hdl.UpdateStaysReports)
	router.Get("/report", hdl.GetAllStaysReports)

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("should be delete error", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodDelete, "/report/"+testUserID.String(), nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("DeleteStaysReports", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("failed to delete stay report"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
	t.Run("should be get error", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/report", nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("GetAllStaysReports", mock.Anything, mock.Anything).Return(nil, errors.New("failed to fetch reports"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestStaysReportsHandler_Create_ParamsError(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.StaysReportsService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	router.Use(handler.JWTMiddleware("your-secret-key", log))

	router.Post("/report/create/{stayId}", hdl.CreateStaysReports)
	router.Delete("/report/{reportId}", hdl.DeleteStaysReports)
	router.Patch("/report/{reportId}", hdl.UpdateStaysReports)
	router.Get("/report", hdl.GetAllStaysReports)

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("should return error if image parsing fails post", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPost, "/report/create/"+testUserID.String(), nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		req.Header.Set("Content-Type", "multipart/form-data")
		req.Body = ioutil.NopCloser(bytes.NewReader([]byte("invalid content")))

		svc.On("CreateStaysReports", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("erorr"))

		router.ServeHTTP(r, req)
		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
	t.Run("should return error if image parsing fails patch", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPatch, "/report/"+testUserID.String(), nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		req.Header.Set("Content-Type", "multipart/form-data")
		req.Body = ioutil.NopCloser(bytes.NewReader([]byte("invalid content")))

		svc.On("UpdateStaysReports", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("erorr"))

		router.ServeHTTP(r, req)
		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
	t.Run("should return error when no image is provided", func(t *testing.T) {
		r := httptest.NewRecorder()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		writer.WriteField("title", "Test Title")
		writer.WriteField("description", "Test Description")

		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/report/create/"+testUserID.String(), body)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		req.Header.Set("Content-Type", "multipart/form-data; boundary="+writer.Boundary())

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
	t.Run("should return error when no image is provided patch", func(t *testing.T) {
		r := httptest.NewRecorder()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		boundary := writer.Boundary()

		writer.Close()

		req := httptest.NewRequest(http.MethodPatch, "/report/"+testUserID.String(), body)
		req.Header.Set("Content-Type", "multipart/form-data; boundary="+boundary)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
	t.Run("should be error with the image type", func(t *testing.T) {
		r := httptest.NewRecorder()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		boundary := writer.Boundary()

		partHeader := make(textproto.MIMEHeader)
		partHeader.Set("Content-Type", "image/svg")
		partHeader.Set("Content-Disposition", `form-data; name="image"; filename="test.svg"`)
		part, err := writer.CreatePart(partHeader)
		if err != nil {
			t.Fatal(err)
		}

		jpegContent := []byte{
			0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46, 0x00, 0x01,
			0x01, 0x01, 0x00, 0x60, 0x00, 0x60, 0x00, 0x00, 0xFF, 0xD9,
		}
		part.Write(jpegContent)

		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/report/create/"+testUserID.String(), body)
		req.Header.Set("Content-Type", "multipart/form-data; boundary="+boundary)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
	t.Run("should be error with the image type patch", func(t *testing.T) {
		r := httptest.NewRecorder()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		boundary := writer.Boundary()

		partHeader := make(textproto.MIMEHeader)
		partHeader.Set("Content-Type", "image/svg")
		partHeader.Set("Content-Disposition", `form-data; name="image"; filename="test.svg"`)
		part, err := writer.CreatePart(partHeader)
		if err != nil {
			t.Fatal(err)
		}

		jpegContent := []byte{
			0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46, 0x00, 0x01,
			0x01, 0x01, 0x00, 0x60, 0x00, 0x60, 0x00, 0x00, 0xFF, 0xD9,
		}
		part.Write(jpegContent)

		writer.Close()

		req := httptest.NewRequest(http.MethodPatch, "/report/"+testUserID.String(), body)
		req.Header.Set("Content-Type", "multipart/form-data; boundary="+boundary)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
	t.Run("should be no errors with jpeg", func(t *testing.T) {
		r := httptest.NewRecorder()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		boundary := writer.Boundary()

		partHeader := make(textproto.MIMEHeader)
		partHeader.Set("Content-Type", "image/jpeg")
		partHeader.Set("Content-Disposition", `form-data; name="image"; filename="test.svg"`)
		part, err := writer.CreatePart(partHeader)
		if err != nil {
			t.Fatal(err)
		}

		jpegContent := []byte{
			0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46, 0x00, 0x01,
			0x01, 0x01, 0x00, 0x60, 0x00, 0x60, 0x00, 0x00, 0xFF, 0xD9,
		}
		part.Write(jpegContent)

		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/report/create/"+testUserID.String(), body)
		req.Header.Set("Content-Type", "multipart/form-data; boundary="+boundary)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
	t.Run("should be no errors with jpeg patch", func(t *testing.T) {
		r := httptest.NewRecorder()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		boundary := writer.Boundary()

		partHeader := make(textproto.MIMEHeader)
		partHeader.Set("Content-Type", "image/jpeg")
		partHeader.Set("Content-Disposition", `form-data; name="image"; filename="test.svg"`)
		part, err := writer.CreatePart(partHeader)
		if err != nil {
			t.Fatal(err)
		}

		jpegContent := []byte{
			0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46, 0x00, 0x01,
			0x01, 0x01, 0x00, 0x60, 0x00, 0x60, 0x00, 0x00, 0xFF, 0xD9,
		}
		part.Write(jpegContent)

		writer.Close()

		req := httptest.NewRequest(http.MethodPatch, "/report/"+testUserID.String(), body)
		req.Header.Set("Content-Type", "multipart/form-data; boundary="+boundary)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
	t.Run("should be errors with jpeg content", func(t *testing.T) {
		r := httptest.NewRecorder()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		boundary := writer.Boundary()

		partHeader := make(textproto.MIMEHeader)
		partHeader.Set("Content-Type", "image/jpeg")
		partHeader.Set("Content-Disposition", `form-data; name="image"; filename="test.svg"`)
		part, err := writer.CreatePart(partHeader)
		if err != nil {
			t.Fatal(err)
		}

		// incorrect image type
		jpegContent := []byte{}
		part.Write(jpegContent)

		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/report/create/"+testUserID.String(), body)
		req.Header.Set("Content-Type", "multipart/form-data; boundary="+boundary)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
	t.Run("should be svc error creating", func(t *testing.T) {
		r := httptest.NewRecorder()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		boundary := writer.Boundary()

		partHeader := make(textproto.MIMEHeader)
		partHeader.Set("Content-Type", "image/jpeg")
		partHeader.Set("Content-Disposition", `form-data; name="image"; filename="test.jpg"`)
		part, err := writer.CreatePart(partHeader)
		if err != nil {
			t.Fatal(err)
		}

		jpegContent := []byte{
			0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46, 0x00, 0x01,
			0x01, 0x01, 0x00, 0x60, 0x00, 0x60, 0x00, 0x00, 0xFF, 0xD9,
		}
		part.Write(jpegContent)

		writer.WriteField("title", "Test Title")
		writer.WriteField("description", "Test Description")

		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/report/create/"+testUserID.String(), body)
		req.Header.Set("Content-Type", "multipart/form-data; boundary="+boundary)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("CreateStaysReports", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(errors.New("svc error"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})

}

func TestStaysReportsHandler_Create_ParamsError_Patch(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.StaysReportsService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	router.Use(handler.JWTMiddleware("your-secret-key", log))

	router.Patch("/report/{reportId}", hdl.UpdateStaysReports)

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("should return error with invalid jpeg content in patch error", func(t *testing.T) {
		r := httptest.NewRecorder()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		boundary := writer.Boundary()

		partHeader := make(textproto.MIMEHeader)
		partHeader.Set("Content-Type", "image/jpeg")
		partHeader.Set("Content-Disposition", `form-data; name="image"; filename="test.jpg"`)
		part, err := writer.CreatePart(partHeader)
		if err != nil {
			t.Fatal(err)
		}

		jpegContent := []byte{
			0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46, 0x00, 0x01,
			0x01, 0x01, 0x00, 0x60, 0x00, 0x60, 0x00, 0x00, 0xFF, 0xD9,
		}
		part.Write(jpegContent)

		writer.WriteField("title", "Test Title")
		writer.WriteField("description", "Test Description")

		writer.Close()

		req := httptest.NewRequest(http.MethodPatch, "/report/"+testUserID.String(), body)
		req.Header.Set("Content-Type", "multipart/form-data; boundary="+boundary)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("UpdateStaysReports", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errors.New("svc error")) // Возвращаем пустой объект и nil (если это нужно)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestStaysReportsHandler_GetStaysById_ErrNoRows(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.StaysReportsService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	router.Use(handler.JWTMiddleware("your-secret-key", log))

	router.Get("/report/{stayId}", hdl.GetStaysReportById)

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("get report with report not found", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/report/"+testUserID.String(), nil)

		svc.On("GetStaysReportById", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, sql.ErrNoRows)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusNotFound, r.Code)
	})
}

func TestStaysReportsHandler_GetStaysById_Error(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.StaysReportsService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	router.Use(handler.JWTMiddleware("your-secret-key", log))

	router.Get("/report/{stayId}", hdl.GetStaysReportById)

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("get report with report not found", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/report/"+testUserID.String(), nil)

		svc.On("GetStaysReportById", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errors.New("failed to fetch report"))

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusNotFound, r.Code)
	})

}

func TestStaysReportsHandler_Create_ParamsError_Patch_Success(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.StaysReportsService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	router.Use(handler.JWTMiddleware("your-secret-key", log))

	router.Post("/report/create/{stayId}", hdl.CreateStaysReports)
	router.Delete("/report/{reportId}", hdl.DeleteStaysReports)
	router.Patch("/report/{reportId}", hdl.UpdateStaysReports)
	router.Get("/report", hdl.GetAllStaysReports)

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("should return error with invalid jpeg content in patch error", func(t *testing.T) {
		r := httptest.NewRecorder()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		boundary := writer.Boundary()

		partHeader := make(textproto.MIMEHeader)
		partHeader.Set("Content-Type", "image/jpeg")
		partHeader.Set("Content-Disposition", `form-data; name="image"; filename="test.jpg"`)
		part, err := writer.CreatePart(partHeader)
		if err != nil {
			t.Fatal(err)
		}

		jpegContent := []byte{
			0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46, 0x00, 0x01,
			0x01, 0x01, 0x00, 0x60, 0x00, 0x60, 0x00, 0x00, 0xFF, 0xD9,
		}
		part.Write(jpegContent)

		writer.WriteField("title", "Test Title")
		writer.WriteField("description", "Test Description")

		writer.Close()

		req := httptest.NewRequest(http.MethodPatch, "/report/"+testUserID.String(), body)
		req.Header.Set("Content-Type", "multipart/form-data; boundary="+boundary)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("UpdateStaysReports", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})
}

func TestStaysReportsHandler_ReportCreateSuccess(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.StaysReportsService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()
	router.Use(handler.JWTMiddleware("your-secret-key", log))
	router.Post("/report/create/{stayId}", hdl.CreateStaysReports)
	router.Delete("/report/{reportId}", hdl.DeleteStaysReports)
	router.Patch("/report/{reportId}", hdl.UpdateStaysReports)
	router.Get("/report", hdl.GetAllStaysReports)
	router.Get("/report/{stayId}", hdl.GetStaysReportById)

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("should be delete success", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodDelete, "/report/"+testUserID.String(), nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("DeleteStaysReports", mock.Anything, mock.Anything, mock.Anything).Return(nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})
	t.Run("should be get all success", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/report", nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("GetAllStaysReports", mock.Anything, mock.Anything).Return(nil, nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)

	})
	t.Run("should be get one success", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/report/"+testUserID.String(), nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("GetStaysReportById", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)

	})
	t.Run("should be success updating a new report", func(t *testing.T) {
		r := httptest.NewRecorder()

		payload := `{"title": "smth", "description": "smth"}`

		req := httptest.NewRequest(http.MethodPatch, "/report/"+testUserID.String(), strings.NewReader(payload))
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("UpdateStaysReports", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
	t.Run("should be no errors creating", func(t *testing.T) {
		r := httptest.NewRecorder()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		boundary := writer.Boundary()

		partHeader := make(textproto.MIMEHeader)
		partHeader.Set("Content-Type", "image/jpeg")
		partHeader.Set("Content-Disposition", `form-data; name="image"; filename="test.jpg"`)
		part, err := writer.CreatePart(partHeader)
		if err != nil {
			t.Fatal(err)
		}

		jpegContent := []byte{
			0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46, 0x00, 0x01,
			0x01, 0x01, 0x00, 0x60, 0x00, 0x60, 0x00, 0x00, 0xFF, 0xD9,
		}
		part.Write(jpegContent)

		writer.WriteField("title", "Test Title")
		writer.WriteField("description", "Test Description")

		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/report/create/"+testUserID.String(), body)
		req.Header.Set("Content-Type", "multipart/form-data; boundary="+boundary)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("CreateStaysReports", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusCreated, r.Code)
	})

}
