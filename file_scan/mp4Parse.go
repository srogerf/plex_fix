package file_scan

import (
	"log"
	"os"
)

func GuessMp4(file os.FileInfo, path string) {
	log.Printf("its a file: %v\n", file.Name())
	ScanTree(path + "/" + file.Name())
}
