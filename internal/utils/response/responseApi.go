package responseApi

import (
	"github.com/go-chi/render"
	"net/http"
)

type ResponseApi struct {
	Status   http.ConnState `json:"status"`
	Response interface{}    `json:"response"`
}

func WriteJson(w http.ResponseWriter, r *http.Request, status http.ConnState, data interface{}) {
	render.Status(r, int(status))
	render.JSON(w, r, &ResponseApi{
		Status:   status,
		Response: data,
	})
}
