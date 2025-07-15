package usecase

import (
	"context"
	"errors"
	pb "github.com/f0xdl/file-processor-grpc/gen/go/fileprocessor"
	"github.com/f0xdl/file-processor-grpc/internal/domain"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"sync"
)

const MaxJobs = 5

type IFileProcessor interface {
	FileExist(ctx context.Context, path string) bool
	GetStats(ctx context.Context, path string) domain.FileStats
}

type FileProcessorUC struct {
	pb.UnimplementedFileProcessorServer
	store IFileProcessor
}

func NewFileServiceServer(store IFileProcessor) *FileProcessorUC {
	return &FileProcessorUC{store: store}
}

func (p *FileProcessorUC) GetFileStats(list *pb.FileList, g grpc.ServerStreamingServer[pb.FileStats]) error {
	if list == nil {
		return errors.New("list is nil")
	}
	fileCount := len(list.Paths)
	if fileCount == 0 {
		return nil
	}

	//prepare fan-in/out
	wg := sync.WaitGroup{}
	sem := make(chan struct{}, MaxJobs)
	results := make(chan domain.FileStats, fileCount)

	// fan-out
	ctx := g.Context()
	for id, path := range list.Paths {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()

			sem <- struct{}{}
			defer func() {
				<-sem
			}()

			if ctx.Err() != nil {
				results <- domain.FileStats{Path: path, Err: ctx.Err()}
				return
			}

			log.Info().Int("job", id).Str("path", path).Msg("file processing start")
			results <- p.store.GetStats(ctx, path)
			log.Info().Int("job", id).Str("path", path).Msg("file processing done")
		}(path)
	}

	//close channels
	go func() {
		wg.Wait()
		close(results)
	}()

	// fan-in
	for fileStats := range results {
		if err := g.SendMsg(FileStatsDtoPb(fileStats)); err != nil {
			log.Error().Err(err).Str("path", fileStats.Path).Msg("send error")
		} else {
			log.Info().Str("path", fileStats.Path).Msg("send ok")
		}
	}
	return nil
}
