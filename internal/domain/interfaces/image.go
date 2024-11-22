package interfaces

import (
	"net/http"
)

type ImageHandler interface {
	UploadImage(w http.ResponseWriter, r *http.Request)
}
