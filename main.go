package main

import (
	"fmt"
	"log"

	download "github.com/ixxiv/download-manager/donwload"
)

func main() {
	download, err := download.Download("", "", 10) // test
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(download)
}
