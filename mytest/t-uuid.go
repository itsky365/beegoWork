package main

import (
    "github.com/sluu99/uuid"
    "fmt"
)

func main() {
    id := uuid.Rand()
    fmt.Println(id.Hex())
    //fmt.Println(id.Raw())


    id1, err := uuid.FromStr("1870747d-b26c-4507-9518-1ca62bc66e5d")
    if err != nil {
        fmt.Println(err)
    }
    id2 := uuid.MustFromStr("1870747db26c450795181ca62bc66e5d")
    fmt.Println(id2)
    fmt.Println(id1 == id2) // true
}
