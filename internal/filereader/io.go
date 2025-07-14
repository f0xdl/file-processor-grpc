package filereader

import (
	"bufio"
	"context"
	"github.com/f0xdl/file-processor-grpc/internal/domain"
	"os"
	"path/filepath"
	"strings"
)

type IoFileReader struct {
	basePath string
}

func NewIoFileReader(basePath string) *IoFileReader {
	return &IoFileReader{basePath: basePath}
}

func (i IoFileReader) FileExist(_ context.Context, path string) bool {
	path = filepath.Join(i.basePath, path)
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

func (i IoFileReader) GetStats(ctx context.Context, path string) (s domain.FileStats) {
	s.Path = path
	file, err := os.Open(filepath.Join(i.basePath, path))
	if err != nil {
		s.Err = err
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			s.Err = ctx.Err()
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
