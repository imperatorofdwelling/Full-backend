package repo

import "errors"

var (
	ErrNotFound     = errors.New("not found")
	ErrUpdateFailed = errors.New("update failed")
)
