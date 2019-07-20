package file_scan

import (
	"os"
	"path/filepath"
)

type Meta struct {
	name      string
	videoType string
	Show      string
	episode   int
	season    string
}

func GuessMetadata(file os.FileInfo, path string) *Meta {
	ext := filepath.Ext(file.Name())
	switch ext {
	case `.mp4`:
		return GuessMp4(file, path)
	default:
		//log.Printf("no scanner %s type %v\n", file.Name(), ext)
		return nil
	}
}
