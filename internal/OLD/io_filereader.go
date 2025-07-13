package OLD

import (
	"github.com/f0xdl/file-processor-grpc/internal/domain"
	"os"
	"path/filepath"
)

type IoFileReader struct {
	basePath string
}

func NewIoFileReader(basePath string) *IoFileReader {
	return &IoFileReader{basePath: basePath}
}

func (fr *IoFileReader) GetBytesFromFile(path string) ([]byte, error) {
	if path == "" {
		return nil, domain.WrongFileNameError(path)
	}
	absPath, err := filepath.Abs(filepath.Join(fr.basePath, path))
	if err != nil {
		return nil, err
	}

	return os.ReadFile(absPath)
}
