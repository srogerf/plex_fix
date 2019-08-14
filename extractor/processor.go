package extractor

import (
	"github.com/srogerf/plex_fix/data"
	"log"
)

//process file given the metadata
func Process(meta *data.MediaData) {
	if meta == nil {
		return
	}

	//figure out correct name
	log.Printf("Series::::%s\n", meta.Show)
	log.Printf("Name::::%s\n", meta.Name)
}
