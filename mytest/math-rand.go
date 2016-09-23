package main

import (
    "math/rand"
    "fmt"
    "time"
    //"strings"
    "strconv"
    "strings"
)

func SmsCode() string {
    var randSlice = make([]string, 0)
    // 根据时间设置随机数种子
    rand.Seed(int64(time.Now().Nanosecond()))
    // 获取指定范围内的随机数
    for i := 0; i < 6; i++ {
        randNum := rand.Intn(10)
        //fmt.Printf("%v ", randNum)
        randSlice = append(randSlice, strconv.Itoa(randNum))
    }
    //fmt.Println(randList)
    ret := strings.Join(randSlice, "")
    //fmt.Println(ret)
    return ret
}

func main() {
    //fmt.Print(rand.Intn(100), ",")
    //fmt.Print(rand.Intn(100))
    //fmt.Println()
    //
    //s1 := rand.NewSource(42)
    //r1 := rand.New(s1)
    //fmt.Println(r1.Int())
    //
    //for i := 0; i < 5; i++ {
    //    fmt.Printf("%v ", rand.Int())
    //}
    //fmt.Println()

    s := SmsCode()
    fmt.Println(s)
}
