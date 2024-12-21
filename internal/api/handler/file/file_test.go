package file

import (
	"github.com/go-chi/chi/v5"
	"github.com/imperatorofdwelling/Full-backend/internal/config"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces/mocks"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger"
	"testing"
)

func TestHandler_NewFileHandler(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.FileService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	t.Run("should be no errors", func(t *testing.T) {
		hdl.NewFileHandler(router)
	})
}
