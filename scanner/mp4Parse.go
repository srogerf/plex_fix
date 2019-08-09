package scanner

import (
	log "github.com/sirupsen/logrus"
	"os"
			"github.com/srogerf/plex_fix/data"
)

func GuessMp4(file os.FileInfo, path string) *data.MediaData {
	log.Printf("its a file: %v\n", file.Name())
	videoMeta := ScanTree(path + "/" + file.Name())
	log.Printf("Video: %s", videoMeta.Show)
	return videoMeta
}
