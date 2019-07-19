package main

import (
	"github.com/srogerf/plex_fix/file_scan"
	"log"
)

func main() {
	//count := file_scan.List("x:/video/downloaded/television")
	count := file_scan.List("/mnt/x/video/downloaded/television")
	log.Printf("Directory has %d items", count)
}
