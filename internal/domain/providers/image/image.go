package image

import (
	"github.com/google/wire"
	"github.com/imperatorofdwelling/Full-backend/internal/api/handler/image"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	createFile "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/file" // Import FileProviderSet
	"log/slog"
)

var ImageProviderSet = wire.NewSet(
	createFile.FileProviderSet,
	ProvideImageHandler,
)

func ProvideImageHandler(svc interfaces.FileService, log *slog.Logger) *image.Handler {
	return &image.Handler{
		Svc: svc,
		Log: log,
	}
}
