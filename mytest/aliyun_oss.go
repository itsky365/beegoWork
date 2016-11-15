package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var Endpoint = "oss-cn-shenzhen.aliyuncs.com"
var AccessKeyId = "***"
var AccessKeySecret = "***"
var BucketName = "***"

func main() {
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

	err = bucket.PutObjectFromFile("LocalFile22.txt", "./good.txt")
	if err != nil {
		fmt.Println(err)
	}
}
