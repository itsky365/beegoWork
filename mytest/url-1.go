package main

import (
    "fmt"
    "net/url"
)

func main() {
    url1 := "http://www.baidu.com/?s=sss&x=xxx#b=name"
    u, err := url.Parse(url1)
    if err != nil {
        panic(err)
    }
    fmt.Println(u)
    fmt.Println(u.RawQuery)
    fmt.Println(u.Query())
    v := u.Query()
    fmt.Println(v["s"])

    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m)
    fmt.Println(m["s"][0])
}
