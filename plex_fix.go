/**
  scan a directory for media files and make it conform to plex standards
*/
package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"github.com/srogerf/plex_fix/extractor"
	"github.com/srogerf/plex_fix/scanner"
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

	//scan the files
	count := scanner.List(*path, extractor.Process)

	log.Printf("Directory %s has %d items", *path, count)

}
