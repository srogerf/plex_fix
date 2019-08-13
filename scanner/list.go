package scanner

import (
	log "github.com/sirupsen/logrus"
	"github.com/srogerf/plex_fix/data"
	"io/ioutil"
)

func List(root_dir string, handler data.ProcessMetadata) (count int) {
	log.Println(root_dir)
	files, err := ioutil.ReadDir(root_dir)
	if err != nil {
		log.Fatal(err)
	}
	count, items := 0, 0

	for _, f := range files {
		//log.Printf("scan: >%s\n", f.Name())
		if f.IsDir() {
			items = List(root_dir+"/"+f.Name(), handler)
		} else {
			items = 0
			m_data := GuessMetadata(f, root_dir)
			handler(m_data)
			count++
		}
		count += items
	}
	log.Printf("Counted %d filesi in %s\n", count, root_dir)
	return count
}
