package handler

import "github.com/pkg/errors"

var (
	ErrImageTypeNotSvg  = errors.New("content type is not image/svg+xml")
	ErrInvalidImageSize = errors.New("size must be greater than 0 and less than 1MB")
)
