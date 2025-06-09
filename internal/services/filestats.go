package services

import (
	"bufio"
	"bytes"
	"errors"
	"strings"

	"github.com/f0xdl/file-processor-grpc/internal/domain"
)

type IStorage interface {
	GetBytesFromFile(path string) ([]byte, error)
}

type FileStatsService struct {
	storage IStorage
}

func NewFileStatsService(storage IStorage) *FileStatsService {
	return &FileStatsService{storage: storage}
}

func (fp *FileStatsService) GetFileStats(path string) (stats domain.FileStats, err error) {
	stats.Path = path
	if stats.Path == "" {
		return stats, errors.New("file path cannot be empty")
	}

	data, err := fp.storage.GetBytesFromFile(stats.Path)
	if err != nil {
		return
	} else if len(data) == 0 {
		return
	}

	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		stats.Lines++
		stats.Words += len(strings.Fields(line))
	}

	return stats, scanner.Err()
}
