package file

import (
	"bufio"
	"context"
	"github.com/f0xdl/file-processor-grpc/internal/domain"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const StatsTimeout = time.Minute

type IoFileProcessor struct {
	basePath string
	files    chan string
}

func NewIoFileReader(basePath string) *IoFileProcessor {
	return &IoFileProcessor{
		basePath: basePath,
		files:    make(chan string, 5),
	}
}

func (i *IoFileProcessor) FileExist(_ context.Context, path string) bool {
	path = filepath.Join(i.basePath, path)
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

func (i *IoFileProcessor) GetStats(ctx context.Context, path string) (s *domain.FileStats) {
	s = &domain.FileStats{Path: path}
	ctxTimeout, cancel := context.WithTimeout(ctx, StatsTimeout)
	defer cancel()

	// open file
	file, err := os.Open(filepath.Join(i.basePath, path))
	if err != nil {
		s.Err = err
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		select {
		case <-ctxTimeout.Done():
			s.Err = ctxTimeout.Err()
			return
		default:
		}

		line := scanner.Text()
		s.Lines++
		s.Words += len(strings.Fields(line))
	}

	if err = scanner.Err(); err != nil {
		s.Err = err
	}
	return
}

func (i *IoFileProcessor) SaveFile(ctx context.Context, filename string, content []byte) error {
	err := make(chan error)
	go func() {
		fullPath := filepath.Join(i.basePath, filename)
		log.Warn().Msg("fullPath:" + fullPath)
		err <- os.WriteFile(fullPath, content, 0644)
	}()

	select {
	case result := <-err:
		return result
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (i *IoFileProcessor) StoreExist() bool {
	_, err := os.Stat(i.basePath)
	log.Warn().Str("path", i.basePath).Msg("store path")
	return err == nil
}
