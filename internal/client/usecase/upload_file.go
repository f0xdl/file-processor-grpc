package usecase

import (
	"context"
	"errors"
	"io"
)

type UploadFileUC struct {
}

func NewUploadFileUC() *UploadFileUC {
	return &UploadFileUC{}
}

func (uc *UploadFileUC) Execute(ctx context.Context, name string, r io.Reader) error {
	return errors.New("not implemented")
}
