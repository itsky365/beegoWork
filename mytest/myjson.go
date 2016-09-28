package main

import (
    "encoding/json"
    "fmt"
)

type Result struct {
    Status       int    `json:"status"`
    Message      string `json:"message"`
    ErrorCode    int    `json:"error_code"`
    ErrorMessage string `json:"error_message"`
}

func main() {
    json_str0 := `{"status":0,"message":"success"}`
    json_str1 := `{"status":1,"error_code":5,"error_message":"error"}`

    res0 := Result{}
    res1 := Result{}

    err0 := json.Unmarshal([]byte(json_str0), &res0)
    err1 := json.Unmarshal([]byte(json_str1), &res1)

    fmt.Println(res0, err0)
    fmt.Println(res1, err1)

}