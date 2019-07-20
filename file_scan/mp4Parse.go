package file_scan

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func GuessMp4(file os.FileInfo, path string) *Meta {
	log.Printf("its a file: %v\n", file.Name())
	videoMeta := ScanTree(path + "/" + file.Name())
	log.Printf("Video: %s", videoMeta.Show)
	return videoMeta
}
