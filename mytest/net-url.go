package main

import (
	"fmt"
	"net/url"
	"strconv"
)

// url参数page值重置

func main() {
	myurl := "http://www.baidu.com/aa/bb?q=111&b=bbb&page=11"
	myurl1 := url.QueryEscape(myurl)
	fmt.Println(myurl1)

	fmt.Println(url.QueryUnescape(myurl1))

	page := 10

	link, _ := url.ParseRequestURI(myurl)
	values := link.Query()
	fmt.Println(values)
	if page == 1 {
		values.Del("page")
	} else {
		values.Set("page", strconv.Itoa(page))
	}
	link.RawQuery = values.Encode()
	fmt.Println(link.String())
}
