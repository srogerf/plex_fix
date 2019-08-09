package extractor

import (
    "log"
	"github.com/srogerf/plex_fix/scanner"
)

func Process(meta *scanner.Meta) {
		if meta == nil {
			return
		}
		log.Printf("Series::::%s\n", meta.Show)
}
