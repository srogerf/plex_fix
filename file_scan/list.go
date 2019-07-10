package file_scan

import (
	"io/ioutil"
	"log"
)

func init() {

}

func List(root_dir string) {
	log.Println(root_dir)
	files, err := ioutil.ReadDir(root_dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		log.Printf("scan: >%s\n", f.Name())
	}
}
