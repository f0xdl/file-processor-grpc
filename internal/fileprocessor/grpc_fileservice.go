package file_processor

import (
	"github.com/f0xdl/file-processor-grpc/internal/domain"
	"github.com/f0xdl/file-processor-grpc/internal/services"
)

type FileStats struct {
	domain.FileStats
	Err error
}

type RpcFileService struct {
	fileProcessor services.FileProcessor
}

func (service *RpcFileService) ProcessFiles(files []string) []FileStats {
	stats := make([]FileStats, len(files))
	// for i, file := range files {
	// 	stats[i].FileStats, stats[i].Err = service.fileProcessor.GetFileStats(file)
	// }
	return stats
}
