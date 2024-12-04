package interfaces

import (
	"github.com/imperatorofdwelling/Full-backend/internal/service/file"
	"net/http"
)

type FileService interface {
	UploadImage(img []byte, imgType file.ImageType, category string) (string, error)
	RemoveFile(fileName string) error
	GenRandomFileName() (string, error)
}

type FileHandler interface {
	GetStaticImage() http.Handler
}
