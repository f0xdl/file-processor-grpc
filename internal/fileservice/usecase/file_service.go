package usecase

import (
	pb "github.com/f0xdl/file-processor-grpc/gen/go/fileprocessor"
	"google.golang.org/grpc"
)

type FileProcessor struct {
	pb.UnimplementedFileProcessorServer
}

func NewFileServiceServer() *FileProcessor {
	return &FileProcessor{}
}

func (f FileProcessor) GetFileStats(list *pb.FileList, g grpc.ServerStreamingServer[pb.FileStats]) error {
	//	TODO implement me
	panic("implement me")
}

//func (f FileProcessor) UploadFile(ctx context.Context, req *pb.UploadFileReq) (*pb.UploadFileResp, error) {
//	//	TODO implement me
//panic("implement me")
//}
