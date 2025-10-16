package usecase

import (
	pb "github.com/f0xdl/file-processor-grpc/api/generated/fileprocessor"
	"github.com/f0xdl/file-processor-grpc/internal/domain"
)

func FileStatsDtoPb(d domain.FileStats) *pb.FileStats {
	return &pb.FileStats{
		Path:  d.Path,
		Lines: int32(d.Lines),
		Words: int32(d.Words),
		Err:   d.Error(),
	}
}
