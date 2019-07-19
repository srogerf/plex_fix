package file_scan

import (
	"io/ioutil"
	"log"
)

func List(root_dir string) (count int) {
	log.Println(root_dir)
	files, err := ioutil.ReadDir(root_dir)
	if err != nil {
		log.Fatal(err)
	}
	count, items := 0, 0

	for _, f := range files {
		//log.Printf("scan: >%s\n", f.Name())
		if f.IsDir() {
			items = List(root_dir + "/" + f.Name())
		} else {
			items = 0
			GuessMetadata(f, root_dir)
			count++
		}
		count += items
	}
	log.Printf("Counted %d filesi in %s\n", count, root_dir)
	return
}
