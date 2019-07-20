// scan the path the file is in for additional metadata
package file_scan

import (
	log "github.com/sirupsen/logrus"
	"path/filepath"
	"strings"
)

// parse given path and figure out show
func ScanTree(path string) *Meta {
	log.Printf("scanning %s\n", path)
	result := new(Meta)
	indepPath := filepath.FromSlash(path)
	dirs := strings.Split(indepPath, "/")
	for dirIdx, dir := range dirs {
		if dir == "television" {
			result.videoType = "television"
			result.show = dirs[dirIdx+1]
			result.season = dirs[dirIdx+2]
			result.name = dirs[dirIdx+3]
		}
	}
	return result
}
