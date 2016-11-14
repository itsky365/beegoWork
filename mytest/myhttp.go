package main

import (
    "net/http"
    "fmt"
    "io/ioutil"
    "net/url"
    //"strings"
    "encoding/json"
)



type JsonType struct {
    Message string `json:"message"`
    Result string `json:"result"`
    JsonTypeData `json:"data"`
}

type JsonTypeData struct {
    ErrorCode string `json:"errorCode"`
    ModuleName string `json:"moduleName"`
    ServiceName string `json:"serviceName"`
    Parameter string `json:"parameter"`
}

func httpDo() {
    client := &http.Client{}

    req, err := http.NewRequest("GET", "http://beta-toe.51awifi.com/exceptionlog/show?id=40918", nil)
    //fmt.Println(&http.Request.RequestURI)
    if err != nil {
        // handle error
    }

    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Set("Cookie", "JSESSIONID_ADMIN=D9129FF8B91DFECD2E74CD172E2B3F76")

    resp, err := client.Do(req)

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }

    //respBody := string((body))
    fmt.Println(string((body)))
    var v1 JsonType
    jsonErr := json.Unmarshal(body, &v1)
    if jsonErr != nil {
        fmt.Println(jsonErr)
    }
    fmt.Println(v1)
    //fmt.Println()
    ////fmt.Println(v1["message"])
    //fmt.Println()
    //fmt.Println(v1.result)
    //fmt.Println()
    fmt.Println(v1.JsonTypeData.ErrorCode)
    fmt.Println(v1.JsonTypeData.ModuleName)
    fmt.Println(v1.JsonTypeData.ServiceName)
    fmt.Println(v1.JsonTypeData.Parameter)
    //jsonData := v1["data"].(map[string]interface{})
    //fmt.Println(jsonData)
    //fmt.Println(v1["data"]["serviceName"])
}

func Test()  {
    resp, err := http.Get("http://toe.51awifi.com")
    if err != nil {
        fmt.Println(err)
    }
    defer resp.Body.Close()

    resp_body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(resp_body))


    u := url.Values{"key": {"Value", "value1"}, "id": {"123"}}
    fmt.Println(u.Encode())
}

func main() {
    //Test()

    httpDo()
}
