// scan the path the file is in for additional metadata
package scanner

import (
	log "github.com/sirupsen/logrus"
	"github.com/srogerf/plex_fix/data"
	"path/filepath"
	"strings"
)

// parse given path and figure out show
func ScanTree(path string) *data.MediaData {
	log.Printf("scanning %s\n", path)
	result := new(data.MediaData)
	indepPath := filepath.FromSlash(path)
	dirs := strings.Split(indepPath, "/")
	for dirIdx, dir := range dirs {
		if dir == "television" {
			left_values := len(dirs) - dirIdx
			result.VideoType = "television"
			result.Show = dirs[dirIdx+1]
			result.Season = dirs[dirIdx+2]
			if left_values > 3 {
				result.Name = dirs[dirIdx+3]
			}
		}
	}
	return result
}
