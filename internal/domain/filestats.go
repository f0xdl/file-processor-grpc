package domain

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/rs/zerolog/log"
	"os"
)

type FileStats struct {
	Path  string `json:"path"`
	Lines int    `json:"lines"`
	Words int    `json:"words"`
	Err   error  `json:"err"`
}

func (fs FileStats) MarshalJSON() ([]byte, error) {
	type FStats FileStats
	return json.Marshal(&struct {
		*FStats
		Err string `json:"err,omitempty"`
	}{
		Err:    fs.Err.Error(),
		FStats: (*FStats)(&fs),
	})
}

func (fs FileStats) Error() string {
	switch {
	case fs.Err == nil:
		return ""
	case errors.Is(fs.Err, os.ErrNotExist):
		return "not found"
	case errors.Is(fs.Err, os.ErrPermission):
		return "access denied"
	case errors.Is(fs.Err, context.DeadlineExceeded) || errors.Is(fs.Err, context.Canceled):
		return "timeout"
	default:
		log.Error().Err(fs.Err).Str("path", fs.Path).Msg("unknown error when  opening the file")
		return "unknown error"
	}
}
