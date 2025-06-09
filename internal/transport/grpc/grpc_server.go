package grpc

import (
	"errors"
	"log"

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
	for path := range in.Paths {
		log.Println(path)
	}
	out.Send(&pb.FileStats{Path: in.Paths[0], Error: "test"})

	return errors.New("method ProcessFiles not implemented")
}
