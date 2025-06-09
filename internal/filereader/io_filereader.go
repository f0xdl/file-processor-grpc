package filereader

import (
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
		return nil, WrongFileNameError(path)
	}
	absPath, err := filepath.Abs(filepath.Join(fr.basePath, path))
	if err != nil {
		return nil, err
	}

	return os.ReadFile(absPath)
}
