/**
  scan a directory for media files and make it conform to plex standards
*/
package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"github.com/srogerf/plex_fix/file_scan"
)

func main() {
	//read the command line flags
	path := flag.String("dir", "/mnt/x/video/downloaded/television", "directory to scan for media")
	debug := flag.Bool("debug", false, "debug")
	flag.Parse()

	if *debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.ErrorLevel)
	}

	//handler for returning metadata
	processMetadata := func(meta *file_scan.Meta) {
		if meta == nil {
			return
		}
		log.Printf("Series::::%s\n", meta.Show)
	}

	//scan the files
	count := file_scan.List(*path, processMetadata)

	log.Printf("Directory %s has %d items", *path, count)
}
