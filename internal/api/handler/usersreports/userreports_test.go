package usersreports

import (
	"bytes"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	handler "github.com/imperatorofdwelling/Full-backend/internal/api/handler/user"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces/mocks"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"testing"
	"time"
)

func TestUsersReportsHandler_NewUsersReportsHandler(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.UsersReportsService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	t.Run("should be no errors", func(t *testing.T) {
		hdl.NewUsersReportsHandler(router)
	})
}

func TestUsersReportsHandler_UserIdError(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.UsersReportsService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()
	router.Get("/user/report", hdl.GetAllUsersReports)
	router.Get("/user/report/{reportId}", hdl.GetUsersReportById)
	router.Post("/user/report/create/{toBlameId}", hdl.CreateUsersReports)
	router.Patch("/user/report/{reportId}", hdl.UpdateUsersReports)
	router.Delete("/user/report/{reportId}", hdl.DeleteUsersReports)

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("get all user error", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/user/report", nil)
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

		req := httptest.NewRequest(http.MethodGet, "/user/report/"+testUserID.String(), nil)
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

		req := httptest.NewRequest(http.MethodPost, "/user/report/create/"+testUserID.String(), nil)
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

		req := httptest.NewRequest(http.MethodPatch, "/user/report/"+testUserID.String(), nil)
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

		req := httptest.NewRequest(http.MethodDelete, "/user/report/"+testUserID.String(), nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusUnauthorized, r.Code)
	})
}

func TestUsersReportsHandler_ParamsError(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.UsersReportsService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	router.Use(handler.JWTMiddleware("your-secret-key", log))

	router.Get("/user/report", hdl.GetAllUsersReports)
	router.Get("/user/report/{reportId}", hdl.GetUsersReportById)
	router.Post("/user/report/create/{toBlameId}", hdl.CreateUsersReports)
	router.Patch("/user/report/{reportId}", hdl.UpdateUsersReports)
	router.Delete("/user/report/{reportId}", hdl.DeleteUsersReports)

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("should be params errors post", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPost, "/user/report/create/"+testUserID.String(), nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("CreateUsersReports", mock.Anything, mock.Anything).Return(nil, errors.New("failed to fetch reports"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
	t.Run("should be params errors put", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPatch, "/user/report/"+testUserID.String(), nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("UpdateUsersReports", mock.Anything, mock.Anything).Return(nil, errors.New("failed to fetch reports"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
	t.Run("should be params errors delete", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodDelete, "/user/report/"+testUserID.String(), nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("DeleteUsersReports", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("failed to fetch reports"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
	t.Run("should return error when no image is provided post", func(t *testing.T) {
		r := httptest.NewRecorder()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		writer.WriteField("title", "Test Title")
		writer.WriteField("description", "Test Description")

		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/user/report/create/"+testUserID.String(), body)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		req.Header.Set("Content-Type", "multipart/form-data; boundary="+writer.Boundary())

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
	t.Run("should be error with the image type post", func(t *testing.T) {
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

		req := httptest.NewRequest(http.MethodPost, "/user/report/create/"+testUserID.String(), body)
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

		req := httptest.NewRequest(http.MethodPatch, "/user/report/"+testUserID.String(), body)
		req.Header.Set("Content-Type", "multipart/form-data; boundary="+boundary)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
	t.Run("should be no errors with jpeg post", func(t *testing.T) {
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

		req := httptest.NewRequest(http.MethodPost, "/user/report/create/"+testUserID.String(), body)
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

		req := httptest.NewRequest(http.MethodPatch, "/user/report/"+testUserID.String(), body)
		req.Header.Set("Content-Type", "multipart/form-data; boundary="+boundary)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
	t.Run("should be errors with jpeg content post", func(t *testing.T) {
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

		req := httptest.NewRequest(http.MethodPost, "/user/report/create/"+testUserID.String(), body)
		req.Header.Set("Content-Type", "multipart/form-data; boundary="+boundary)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
	t.Run("should be svc error creating post", func(t *testing.T) {
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

		req := httptest.NewRequest(http.MethodPost, "/user/report/create/"+testUserID.String(), body)
		req.Header.Set("Content-Type", "multipart/form-data; boundary="+boundary)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("CreateUsersReports", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(errors.New("svc error"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
	t.Run("should be svc error creating patch", func(t *testing.T) {
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

		req := httptest.NewRequest(http.MethodPatch, "/user/report/"+testUserID.String(), body)
		req.Header.Set("Content-Type", "multipart/form-data; boundary="+boundary)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("UpdateUsersReports", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errors.New("svc error"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
	t.Run("should be svc all get error", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/user/report", nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("GetAllUsersReports", mock.Anything, mock.Anything).Return(nil, errors.New("error while getting all of the info"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
	t.Run("should be svc one get error", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/user/report/"+testUserID.String(), nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("GetUsersReportById", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("error while getting all of the info"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestUsersReportsHandler_Success(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.UsersReportsService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	router := chi.NewRouter()

	router.Use(handler.JWTMiddleware("your-secret-key", log))

	router.Get("/user/report", hdl.GetAllUsersReports)
	router.Get("/user/report/{reportId}", hdl.GetUsersReportById)
	router.Post("/user/report/create/{toBlameId}", hdl.CreateUsersReports)
	router.Patch("/user/report/{reportId}", hdl.UpdateUsersReports)
	router.Delete("/user/report/{reportId}", hdl.DeleteUsersReports)

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("should no svc error creating", func(t *testing.T) {
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

		req := httptest.NewRequest(http.MethodPost, "/user/report/create/"+testUserID.String(), body)
		req.Header.Set("Content-Type", "multipart/form-data; boundary="+boundary)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("CreateUsersReports", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusCreated, r.Code)
	})
	t.Run("should no svc error creating patch", func(t *testing.T) {
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

		req := httptest.NewRequest(http.MethodPatch, "/user/report/"+testUserID.String(), body)
		req.Header.Set("Content-Type", "multipart/form-data; boundary="+boundary)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("UpdateUsersReports", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})
	t.Run("should be delete success", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodDelete, "/user/report/"+testUserID.String(), nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("DeleteUsersReports", mock.Anything, mock.Anything, mock.Anything).Return(nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})
	t.Run("should be get all success", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/user/report", nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("GetAllUsersReports", mock.Anything, mock.Anything).Return(nil, nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})
	t.Run("should be get one success", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/user/report/"+testUserID.String(), nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("GetUsersReportById", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})
}
