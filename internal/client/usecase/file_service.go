package usecase

import (
	"context"
	"errors"
	"github.com/f0xdl/file-processor-grpc/internal/domain"
)

const MaxFileSize = 1024 * 1024 * 4 // Max 4MB //TODO: migrate to stream send

var MaxFileErr = errors.New("maximum file size is 100MB")

type IHandler interface {
	GetFileInfo(ctx context.Context, names []string) ([]domain.FileStats, error)
	UploadFile(ctx context.Context, name string, data []byte) error
}

type FileService struct {
	handler IHandler
}

func NewFileService(handler IHandler) *FileService {
	return &FileService{handler: handler}
}

func (uc *FileService) UploadFile(ctx context.Context, name string, data []byte) error {
	if len(data) > MaxFileSize {
		return MaxFileErr
	}
	return uc.handler.UploadFile(ctx, name, data)
}

func (uc *FileService) GetFileInfo(ctx context.Context, names []string) ([]domain.FileStats, error) {
	return uc.handler.GetFileInfo(ctx, names)
}
