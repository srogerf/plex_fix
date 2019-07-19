// scan the path the file is in for additional metadata
package file_scan

import (
	"log"
	"path/filepath"
	"strings"
)

// parse given path and figure out show
func ScanTree(path string) {
	indepPath := filepath.FromSlash(path)
	dirs := strings.Split(indepPath, "/")
	for dirIdx, dir := range dirs {
		if dir == "television" {
			log.Printf("tv show---> %v\n", dirs[dirIdx+1])
			log.Printf("   season---> %v\n", dirs[dirIdx+2])
			log.Printf("   episode---> %v\n", dirs[dirIdx+3])
		}
	}
}
