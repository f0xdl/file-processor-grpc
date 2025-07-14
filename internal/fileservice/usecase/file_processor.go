package usecase

import (
	"context"
	"errors"
	pb "github.com/f0xdl/file-processor-grpc/gen/go/fileprocessor"
	"github.com/f0xdl/file-processor-grpc/internal/domain"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type IStore interface {
	FileExist(ctx context.Context, path string) bool
	GetStats(ctx context.Context, path string) domain.FileStats
}

type FileProcessor struct {
	pb.UnimplementedFileProcessorServer
	store IStore
}

func NewFileServiceServer(store IStore) *FileProcessor {
	return &FileProcessor{store: store}
}

func (p FileProcessor) GetFileStats(list *pb.FileList, g grpc.ServerStreamingServer[pb.FileStats]) error {
	if list == nil {
		return errors.New("list is nil")
	}

	for _, path := range list.Paths {
		log.Info().Str("path", path).Msg("file")
		filestat := p.store.GetStats(g.Context(), path)
		err := g.SendMsg(FileStatsDtoPb(filestat))
		if err != nil {
			log.Error().Err(err).Str("path", path).Msg("send error")
		}
	}
	return nil
}
