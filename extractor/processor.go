package extractor

import (
    "log"
    		"github.com/srogerf/plex_fix/data"
)

func Process(meta *data.MediaData) {
		if meta == nil {
			return
		}
		log.Printf("Series::::%s\n", meta.Show)
}
