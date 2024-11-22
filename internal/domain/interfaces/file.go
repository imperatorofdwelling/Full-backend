package interfaces

import (
	"github.com/imperatorofdwelling/Full-backend/internal/service/file"
)

type FileService interface {
	UploadImage(img []byte, imgType file.ImageType, category string) (string, error)
	RemoveFile(fileName string) error
	GenRandomFileName() (string, error)
}
