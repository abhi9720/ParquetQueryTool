package utils

import (
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrInvalidRecord  = errors.New("invalid record")
	ErrSearchFailed   = errors.New("search failed")
	ErrUploadFailed   = errors.New("upload failed")
)
