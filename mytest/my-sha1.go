package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"sort"
	"strings"
)

//对字符串进行MD5哈希
func a(data string) string {
	t := md5.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

//对字符串进行SHA1哈希
func b(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func main() {
	var data string = "abc"
	fmt.Printf("MD5 : %s\n", a(data))
	fmt.Printf("SHA1 : %s\n", b(data))

	// 字符串数组排序
	s := []string{
		"zz",
		"gg",
		"bb",
	}
	sort.Strings(s)
	fmt.Println(s)
	fmt.Println(strings.Join(s, "=="))
}
