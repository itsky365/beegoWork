package main

import (
	"fmt"
	"github.com/ltt1987/alidayu"
)

func main() {
	alidayu.AppKey = "***"
	alidayu.AppSecret = "***"
	alidayu.UseHTTP = true // set UseHTTP to true if prefer http over https

	success, resp := alidayu.SendSMS(
		"***",
		"加国留学",
		"SMS_***",
		`{"smscode":"23456"}`,
	)
	fmt.Println("Success:", success)
	fmt.Println(resp)
}
