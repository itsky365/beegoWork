package main

import (
    "github.com/mholt/archiver"
    "fmt"
)

func main() {
    err := archiver.Zip.Open("upload/componentZip.zip", "upload")
    if err != nil {
        fmt.Println(err)
    }
}
