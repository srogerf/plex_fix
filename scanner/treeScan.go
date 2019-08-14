// scan the path the file is in for additional metadata
package scanner

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/srogerf/plex_fix/data"
	"path/filepath"
	_ "regexp"
	"strings"
)

// parse given path and figure out show
func ScanTree(path string) *data.MediaData {
	log.Printf("******* scanning %s\n", path)
	result := new(data.MediaData)
	indepPath := filepath.FromSlash(path)
	dirs := strings.Split(indepPath, "/")
	for dirIdx, dir := range dirs {
		if dir == "television" {
			remaining_values := len(dirs) - dirIdx - 1
			result.VideoType = "television"
			result.Show = dirs[dirIdx+1]
			result.Season = dirs[dirIdx+2]
			if remaining_values >= 3 {
				result.Name = dirs[dirIdx+3]
			} else {
				log.Printf("needs fixing: %s", result.Season)
				result.Name = result.Season
			}
		}
	}
	fmt.Printf(">>************: %s\n\n", result.Show)
	return result
}
