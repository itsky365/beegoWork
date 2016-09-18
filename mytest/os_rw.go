package main

import (
    "os"
    "fmt"
    //"path"
    //"path/filepath"
)

func main() {
    //pwd, err := os.Getwd()
    //if err != nil {
    //    fmt.Println(err)
    //}
    //fmt.Println(pwd)
    //
    //_path, err := filepath.Abs(pwd)
    //if err != nil {
    //    fmt.Println(err)
    //}
    //fmt.Println(_path)
    //
    //file_path := path.Join(path.Dir(_path), "mytest", "good.txt")
    //fmt.Println("file_path=>", file_path)

    file, err := os.Open("./good.txt")
    if err != nil {
        fmt.Println(err)
    }
    file.Close()

    data := make([]byte, 100)
    count, err := file.Read(data)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Printf("read %d bytes: %q\n", count, data[:count])
}
