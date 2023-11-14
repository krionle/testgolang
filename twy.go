package main

import (
	"fmt"
	"net/http"
)

var urls = []string{
	"https://www.baidu.com/",
	"https://www.bilibili.com/",
	"http://www.google.com/",
}

func main() {
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("err:", url, err)
		}
		fmt.Println(url, ":", resp.Status)
	}
}
