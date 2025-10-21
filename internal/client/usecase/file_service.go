package usecase

import (
	"context"
	"github.com/f0xdl/file-processor-grpc/internal/domain"
	"io"
)

type IHandler interface {
	GetFileInfo(ctx context.Context, names []string) ([]domain.FileStats, error)
}

type FileService struct {
	handler IHandler
}

func NewFileService(handler IHandler) *FileService {
	return &FileService{handler: handler}
}

func (uc *FileService) UploadFile(ctx context.Context, name string, r io.Reader) error {
	//TODO	return uc.handler.UploadFile(ctx, name, r)
	return domain.ErrNotImpl
}

func (uc *FileService) GetFileInfo(ctx context.Context, names []string) ([]domain.FileStats, error) {
	return uc.handler.GetFileInfo(ctx, names)
}
