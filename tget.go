package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	mypost()
}
func myget() {
	requrl := "http://www.baidu.com"
	resp, _ := http.Get(requrl)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
func mypost() {
	requrl := "http://www.baidu.com"
	postValue := url.Values{
		"username": {"aa"},
		"address":  {"bb"},
		"subject":  {"cc"},
		"form":     {"dd"},
	}
	body := bytes.NewBufferString(postValue.Encode())
	request, err := http.Post(requrl, "text/html", body)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(request.StatusCode)
	if request.StatusCode == 200 {
		rb, err := io.ReadAll(request.Body)
		if err != nil {
			print(err)
		}
		fmt.Println(string(rb))
	}
}
