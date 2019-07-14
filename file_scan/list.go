package file_scan

import (
	"io/ioutil"
	"log"
)

func init() {

}

func List(root_dir string) int {
	log.Println(root_dir)
	files, err := ioutil.ReadDir(root_dir)
	if err != nil {
		log.Fatal(err)
	}
	count := 0
	items := 0

	for _, f := range files {
		//log.Printf("scan: >%s\n", f.Name())
		if f.IsDir() {
			items = List(root_dir + "/" + f.Name())
		} else {
			items = 0
			log.Printf("its a file: %s\n", f.Name())
			count++
		}
		count += items
	}
	log.Printf("Counted %d filesi in %s\n", count, root_dir)
	return count
}

//guess show name

//get file name
