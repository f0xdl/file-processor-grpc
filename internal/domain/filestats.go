package domain

import "encoding/json"

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
