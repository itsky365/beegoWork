package main

import(
    "fmt"
    "time"
    "math/rand"
)

const (
    KC_RAND_KIND_NUM   = 0  // 纯数字
    KC_RAND_KIND_LOWER = 1  // 小写字母
    KC_RAND_KIND_UPPER = 2  // 大写字母
    KC_RAND_KIND_ALL   = 3  // 数字、大小写字母
)

// 随机字符串
func Krand(size int, kind int) []byte {
    kindConst := kind
    kinds := [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}
    result := make([]byte, size)

    is_all := kind > 2 || kind < 0
    rand.Seed(time.Now().UnixNano())
    for i :=0; i < size; i++ {
        if is_all { // random ikind
            kindConst = rand.Intn(3)
        }
        //fmt.Println(kindConst)
        scope, base := kinds[kindConst][0], kinds[kindConst][1]
        r := uint8(base + rand.Intn(scope))
        fmt.Println(scope, base, r, string(r))
        result[i] = r
    }
    return result
}

func main(){
    fmt.Println("num:   " + string(Krand(8, KC_RAND_KIND_NUM)))
    fmt.Println("lower: " + string(Krand(8, KC_RAND_KIND_LOWER)))
    fmt.Println("upper: " + string(Krand(8, KC_RAND_KIND_UPPER)))
    fmt.Println("all:   " + string(Krand(8, KC_RAND_KIND_ALL)))
}