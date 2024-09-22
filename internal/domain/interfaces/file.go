package interfaces

import (
	service "github.com/imperatorofdwelling/Website-backend/internal/service/file"
)

type FileService interface {
	UploadImage([]byte, service.ImageType) (string, error)
	RemoveFile(string) error
	GenRandomFileName() (string, error)
}
