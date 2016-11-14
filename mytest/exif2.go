package main

import (
  "fmt"
  "github.com/xiam/exif"
)

func main() {
    data, err := exif.Read("./test1479102187.jpg")
    if err != nil {
        fmt.Println(err)
        return
    }

    for key, val := range data.Tags {
        fmt.Printf("%s = %s\n", key, val)
    }
}