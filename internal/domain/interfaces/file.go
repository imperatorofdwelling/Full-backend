package interfaces

import (
	"github.com/imperatorofdwelling/Full-backend/internal/service/file"
	"net/http"
)

//go:generate mockery --name FileService
type FileService interface {
	UploadImage(img []byte, imgType file.ImageType, filePath file.PathType) (string, error)
	RemoveFile(fileName string) error
	GenRandomFileName() (string, error)
}

type FileHandler interface {
	GetStaticImage() http.Handler
}
