package file_scan

import (
	"log"
	"os"
	"path/filepath"
)

func GuessMetadata(file os.FileInfo, path string) {
	ext := filepath.Ext(file.Name())
	switch ext {
	case `.mp4`:
		GuessMp4(file, path)
	default:
		log.Printf("no scanner %s type %v\n", file.Name(), ext)
	}
}
