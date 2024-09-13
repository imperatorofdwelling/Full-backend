package handler

import (
	services "github.com/imperatorofdwelling/Website-backend/internal/service/interface"
)

type Handler struct {
	service services.ServiceUseCase
}

func NewHandler(service services.ServiceUseCase) *Handler {
	return &Handler{
		service: service,
	}
}
