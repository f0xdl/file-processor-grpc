package grpc

import (
	pb "github.com/f0xdl/file-processor-grpc/gen/go/fileprocessor"
	"github.com/f0xdl/file-processor-grpc/internal/services"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedFileServiceServer
	service *services.FileStatsService
}

func NewRpcServer(service *services.FileStatsService) *server {
	return &server{service: service}
}

func (s *server) ProcessFiles(in *pb.FileList, out grpc.ServerStreamingServer[pb.FileStats]) error {
	for _, path := range in.Paths {
		stats, err := s.service.GetFileStats(path)
		pbStats := pb.FileStats{
			Path:  path,
			Lines: int32(stats.Lines),
			Words: int32(stats.Words),
		}
		if err != nil {
			pbStats.Error = err.Error()
		}
		out.Send(&pbStats)

	}
	return nil
}
