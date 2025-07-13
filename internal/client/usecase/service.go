package usecase

import (
	"context"
	"github.com/f0xdl/file-processor-grpc/internal/domain"
	"io"
)

type FileService struct {
	fileInfo   *GetFileInfoUC
	uploadFile *UploadFileUC
}

func NewFileService(fileInfo *GetFileInfoUC, uploadFile *UploadFileUC) *FileService {
	return &FileService{fileInfo: fileInfo, uploadFile: uploadFile}
}

func (uc *FileService) GetFileInfo(ctx context.Context, names []string) ([]domain.FileStats, error) {
	return uc.fileInfo.ManyExecute(ctx, names)
}

func (uc *FileService) UploadFile(ctx context.Context, name string, r io.Reader) error {
	return uc.uploadFile.Execute(ctx, name, r)
}
