package responseApi

import (
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type ResponseSuccess struct {
	Data interface{} `json:"data"`
}

type ResponseError struct {
	Error string `json:"error"`
}

func WriteJson(w http.ResponseWriter, r *http.Request, status http.ConnState, data interface{}) {
	render.Status(r, int(status))
	render.JSON(w, r,
		&ResponseSuccess{
			Data: data,
		})
}

func WriteError(w http.ResponseWriter, r *http.Request, status http.ConnState, err slog.Attr) {
	render.Status(r, int(status))
	render.JSON(w, r,
		&ResponseError{
			Error: err.String(),
		})
}
