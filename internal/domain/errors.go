package domain

import (
	"errors"
)

var (
	ErrPathNotFound       = errors.New("path not found")
	ErrStoreAccess        = errors.New("store access denied")
	ErrWrongFileName      = errors.New("wrong file name")
	ErrNotImpl            = errors.New("not implemented")
	ErrInternal           = errors.New("internal error")
	ErrServiceUnavailable = errors.New("service unavailable")
	ErrCanceled           = errors.New("canceled")
	ErrTimeout            = errors.New("timeout")
	ErrFileAlreadyExist   = errors.New("file already exists")
)
