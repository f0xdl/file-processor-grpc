package usecase

import (
	"context"
	"errors"
	"github.com/f0xdl/file-processor-grpc/internal/domain"
)

type GetFileInfoUC struct {
}

func NewGetFileInfoUC() *GetFileInfoUC {
	return &GetFileInfoUC{}
}

func (uc *GetFileInfoUC) Execute(ctx context.Context, name string) (domain.FileStats, error) {
	return domain.FileStats{Name: name, Err: errors.New("not implemented")}, nil
}

func (uc *GetFileInfoUC) ManyExecute(ctx context.Context, names []string) ([]domain.FileStats, error) {
	results := make([]domain.FileStats, len(names))
	for i, name := range names {
		results[i], _ = uc.Execute(ctx, name)
	}
	return results, nil
}
