package scanner

import (
	"os"
	"path/filepath"
		"github.com/srogerf/plex_fix/data"
)

func GuessMetadata(file os.FileInfo, path string) *data.MediaData {
	ext := filepath.Ext(file.Name())
	switch ext {
	case `.mp4`:
		return GuessMp4(file, path)
	default:
		//log.Printf("no scanner %s type %v\n", file.Name(), ext)
		return nil
	}
}
