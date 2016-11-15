package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	client := &http.Client{}
	post_dat := url.Values{"username": {"superadmin"}, "password": {"123456"}}
	req, err := http.NewRequest("POST", "http://beta-toe.51awifi.com/login_check", strings.NewReader(post_dat.Encode()))
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	c := resp.Cookies()

	fmt.Println(string(body))

	fmt.Println(c[0].Name)
	fmt.Println(c[0].Value)
	fmt.Println(resp.Cookies())
	fmt.Println(resp.StatusCode)

}
