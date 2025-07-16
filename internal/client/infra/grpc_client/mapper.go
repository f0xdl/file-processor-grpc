package grpc_client

import (
	"errors"
	pb "github.com/f0xdl/file-processor-grpc/gen/go/fileprocessor"
	"github.com/f0xdl/file-processor-grpc/internal/domain"
)

func fileStatsPbToD(v *pb.FileStats) domain.FileStats {
	r := domain.FileStats{
		Path:  v.Path,
		Lines: int(v.Lines),
		Words: int(v.Words),
	}
	if v.Err != "" {
		r.Err = errors.New(v.Err)
	}
	return r
}
