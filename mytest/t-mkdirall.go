package main

import (
	"fmt"
	"io/ioutil"
	"os"
	//"path/filepath"
	"strings"
)

func walkFunc(path string, info os.FileInfo, err error) error {
	fmt.Printf("%s\n", path)
	return nil
}

func main() {
	//err := os.MkdirAll("aa/bb/cc", 0777)
	//if err != nil {
	//    fmt.Println("mkdir err", err)
	//} else {
	//    fmt.Println("mkdir ok")
	//}
	//
	//dir, err := ioutil.ReadDir("aa/bb/cc")
	//fmt.Println(dir)
	//for i, k := range dir {
	//    fmt.Println(i, k, k.Name())
	//}
	//
	////遍历打印所有的文件名
	//filepath.Walk("aa", walkFunc)

	file, err := ioutil.ReadFile("aa/bb/cc/cc1.txt")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(file)
	//fmt.Println(string(file))

	fileContent := string(file)
	//newStr := strings.Replace(fileContent, "__Entity__", "varName11", 1)
	//fmt.Println(newStr)

	r := strings.NewReplacer("__Entity__", "varName11")
	newContent := r.Replace(fileContent)
	fmt.Println(newContent)

	// 读取cc2文件
	fileSetting, err := ioutil.ReadFile("aa/bb/cc/cc2.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileSettingContent := string(fileSetting)
	// 替换内容
	fileSettingContentReplace := strings.NewReplacer("__Setting__", "settingName22")
	fileSettingContentNew := fileSettingContentReplace.Replace(fileSettingContent)
	fmt.Println(fileSettingContentNew)

	contentAll := newContent + "\r\n" + fileSettingContentNew

	err = ioutil.WriteFile("aa/bb/cc/cc12w.txt", []byte(contentAll), 0666)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(err)
}
