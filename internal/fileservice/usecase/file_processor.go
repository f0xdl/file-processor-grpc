package usecase

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"sync"

	pb "github.com/f0xdl/file-processor-grpc/api/generated/fileprocessor"
	"github.com/f0xdl/file-processor-grpc/internal/domain"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

const MaxJobs = 5

type IFileProcessor interface {
	FileExist(ctx context.Context, path string) bool
	GetStats(ctx context.Context, path string) *domain.FileStats
	SaveFile(ctx context.Context, filename string, content []byte) error
	StoreExist() bool
	CalcHash(filename string) ([]byte, error)
}
type ICache interface {
	Get(ctx context.Context, path string) (*domain.FileStats, error)
	Set(ctx context.Context, stats *domain.FileStats) error
}

type FileProcessorUC struct {
	pb.UnimplementedFileProcessorServer
	store IFileProcessor
	cache ICache
}

func NewFileServiceServer(store IFileProcessor, cache ICache) *FileProcessorUC {
	return &FileProcessorUC{
		store: store,
		cache: cache,
	}
}

func (p *FileProcessorUC) GetFileStats(list *pb.FileList, g grpc.ServerStreamingServer[pb.FileStats]) error {
	if list == nil {
		return errors.New("list is nil")
	}
	fileCount := len(list.Paths)
	if fileCount == 0 {
		return nil
	}

	//prepare
	wg := sync.WaitGroup{}
	sem := make(chan struct{}, MaxJobs)
	results := make(chan domain.FileStats, fileCount)

	// fan-out
	ctx := g.Context()
	for id, path := range list.Paths {
		if len(path) == 0 {
			results <- domain.FileStatsError(path, domain.ErrWrongFileName)
		}

		wg.Add(1)
		go func(path string) {
			defer wg.Done()

			sem <- struct{}{}
			defer func() {
				<-sem
			}()
			log.Debug().Int("job", id).Str("path", path).Msg("file processing start")

			if ctx.Err() != nil {
				results <- domain.FileStatsError(path, ctx.Err())
				return
			}

			stats, err := p.cache.Get(ctx, path)
			if err != nil && !errors.Is(err, domain.ErrPathNotFound) {
				log.Warn().Timestamp().Err(err).Str("path", path).Msg("error getting file stats from cache")
			}

			if stats == nil {
				log.Debug().Int("job", id).Str("path", path).
					Msg("file missing in cache, calculating stats")
				stats = p.store.GetStats(ctx, path)
				if stats.Err == nil { // save to cache
					err = p.cache.Set(ctx, stats)
					if err != nil {
						log.Warn().Timestamp().Err(err).Str("path", path).Msg("error set file stats to cache")
					}
				}
			}

			log.Debug().Int("job", id).Str("path", path).Msg("file processing done")
			results <- *stats
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

func (p *FileProcessorUC) IsFileExist(ctx context.Context, r *pb.CheckFileExistsReq) (*wrapperspb.BoolValue, error) {
	if len(r.Filename) == 0 {
		return nil, domain.FileStatsError(r.Filename, domain.ErrWrongFileName)
	}
	return &wrapperspb.BoolValue{
		Value: p.store.FileExist(ctx, r.Filename),
	}, nil
}

func (p *FileProcessorUC) UploadFile(stream grpc.ClientStreamingServer[pb.UploadFileReq, pb.UploadFileRes]) error {
	ctx, cancel := context.WithCancel(stream.Context())
	defer cancel()

	buf := bytes.NewBuffer([]byte{})
	var filename string
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		if req == nil {
			return errors.New("nil input")
		}
		if len(req.Content) == 0 { //stream finished
			break
		}
		if filename == "" {
			filename = req.Filename
			if len(filename) == 0 {
				return domain.FileStatsError(req.Filename, domain.ErrWrongFileName)
			}
		}
		buf.Write(req.Content)
	}

	if !p.store.StoreExist() {
		return domain.FileStatsError(filename, domain.ErrStoreAccess)
	}
	if p.store.FileExist(ctx, filename) {
		return domain.FileStatsError(filename, domain.ErrFileAlreadyExist)
	}
	err := p.store.SaveFile(ctx, filename, buf.Bytes())
	if err != nil {
		return domain.FileStatsError(filename, fmt.Errorf("SaveFile: %w", err))
	}
	hash, err := p.store.CalcHash(filename)
	if err != nil {
		return domain.FileStatsError(filename, fmt.Errorf("CaclHash: %w", err))
	}
	return stream.SendAndClose(&pb.UploadFileRes{Hash: fmt.Sprintf("%x", hash)})
}
