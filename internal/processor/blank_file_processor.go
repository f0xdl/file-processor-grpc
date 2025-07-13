package processor

import "github.com/f0xdl/file-processor-grpc/internal/domain"

type Blank struct {
}

func (b Blank) ProcessFiles(filenames []string) ([]domain.FileStats, error) {
	filestats := make([]domain.FileStats, len(filenames))
	for i, filename := range filenames {
		filestats[i] = domain.FileStats{
			Name:  filename,
			Lines: i,
			Words: i,
		}
	}
	return filestats, nil
}

func NewBlank() *Blank {
	return &Blank{}
}
