package main

import (
    "crypto/sha1"
    "fmt"
)

func main() {
    s := "我是字符串"
    sh := sha1.New()
    sh.Write([]byte(s))
    bs := sh.Sum(nil)

    fmt.Println(s)
    fmt.Printf("%x\n", sh)
    fmt.Printf("%x\n", bs)
}
