package interfaces

import (
	"github.com/imperatorofdwelling/Full-backend/internal/service/file"
)

type FileService interface {
	UploadImage([]byte, file.ImageType, file.ImagePath) (string, error)
	RemoveFile(string) error
	GenRandomFileName() (string, error)
}
