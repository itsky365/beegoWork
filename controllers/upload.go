package controllers

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/logs"
    "github.com/aliyun/aliyun-oss-go-sdk/oss"
    "fmt"
)

var Endpoint = "oss-cn-shenzhen.aliyuncs.com"
var AccessKeyId = "LTAIm4V4AGAK8Il2"
var AccessKeySecret = "DvaBmxgHKeT5vLI8XrSuo0b6U7tuof"
var BucketName = "canadago"

type MyuploadController struct {
    beego.Controller
}

func (c *MyuploadController) Get() {
    c.TplName = "myupload.html"
}


type MyuploadPostController struct {
    beego.Controller
}

func (c *MyuploadPostController) Post() {
    client, err := oss.New(Endpoint, AccessKeyId, AccessKeySecret)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(client)
    
    lsRes, err := client.ListBuckets()
    if err != nil {
        fmt.Println(err)
    }
    
    for _, bucket := range lsRes.Buckets {
        fmt.Println("Buckets:", bucket.Name)
    }
    
    bucket, err := client.Bucket(BucketName)
    if err != nil {
        fmt.Println(err)
    }
    
    
    f, fh, err := c.GetFile("file1")
    if err != nil {
        logs.Debug(err)
    }
    logs.Debug(f, fh)
    c.Ctx.WriteString("upload post")
    //c.SaveToFile("file1", "upload_" + fh.Filename)
    
    // 本地文件上传至oss
    err = bucket.PutObjectFromFile("img22.jpg", "upload_" + fh.Filename)
    if err != nil {
        fmt.Println(err)
    }
    
    // 文件上传至oss
    err = bucket.PutObject("img33.jpg", f)
    if err != nil {
        fmt.Println(err)
    }
    
    return
}
