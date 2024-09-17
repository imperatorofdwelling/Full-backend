package responseApi

import (
	"github.com/go-chi/render"
	"net/http"
)

type ResponseError struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

func WriteJson(w http.ResponseWriter, r *http.Request, status http.ConnState, data interface{}) {
	render.Status(r, int(status))
	render.JSON(w, r, data)
}

func WriteError(w http.ResponseWriter, r *http.Request, status http.ConnState, err error) {
	render.Status(r, int(status))
	render.JSON(w, r, &ResponseError{
		Status: int(status),
		Error:  err.Error(),
	})
}
