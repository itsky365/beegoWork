package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	doc, err := goquery.NewDocument("http://365jia.cn/news/")
	if err != nil {
		log.Fatal(err)
	}

	urlList := []string{}
	fmt.Println(urlList)

	urlChan := make(chan string)
	defer close(urlChan)
	fmt.Println(urlChan)

	// Find the review items
	doc.Find(".tags .item_list").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		title := s.Find("a").Text()
		title = strings.TrimSpace(title)
		href, _ := s.Find("a").Attr("href")
		fmt.Printf("%d: %s => %s\n", i, title, href)
		urlList = append(urlList, href)
	})

	for _, v := range urlList {
		go func() {
			urlChan <- v
		}()
	}

	for k, v := range urlList {
		fmt.Println(v)
		select {
		case res := <-urlChan:
			fmt.Printf("%d => %s\n", k, res)
		}
	}

	fmt.Println(urlList)
	fmt.Println(urlChan)
}

func downContent() {
	// url := "http://img2.bdstatic.com/img/image/166314e251f95cad1c8f496ad547d3e6709c93d5197.jpg"
	// url := "http://news.qq.com/a/20161117/023969.htm"
	url := "http://365jia.cn/news/2016-11-17/DDFDF082CF4104B3.html"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	filename := filepath.Base(url)
	fileExt := filepath.Ext(filename)
	fmt.Println(fileExt)
	fmt.Println(filename)

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	outfile, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	rs, err := io.Copy(outfile, bytes.NewReader(content))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rs)
}

func main() {

	ExampleScrape()
}
