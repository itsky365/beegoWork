package main

import (
	"fmt"
	"github.com/mholt/archiver"
)

func main() {
	err := archiver.Zip.Open("upload/componentZip.zip", "upload")
	if err != nil {
		fmt.Println(err)
	}
}
