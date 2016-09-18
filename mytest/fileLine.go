package main

import (
    "os"
    "fmt"
    "bufio"
)

// 读取文件行数
func main() {
    if len(os.Args) < 2 {
        return
    }
    filename := os.Args[1]
    fmt.Println(filename)

    file, err := os.Open(filename)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    var lineNum int
    for {
        line, prefix, err := reader.ReadLine()
        fmt.Println(string(line))
        fmt.Println(string(line))
        if err != nil {
            break
        }
        if !prefix {
            lineNum++
        }
    }
    fmt.Println(lineNum)
}

// go run fileLine.go file.txt