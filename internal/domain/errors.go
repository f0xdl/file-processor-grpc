package domain

import (
	"errors"
)

var ErrPathNotFound = errors.New("path not found")
var ErrWrongFileName = errors.New("wrong file name")
