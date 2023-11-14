package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func tfile(w http.ResponseWriter, res *http.Request) {
	if res.Method == "GET" {
		data, err := os.ReadFile("tfile.html")
		io.WriteString(w, string(data))
		if err != nil {
			io.WriteString(w, "html 未找到")
			return
		}
	} else if res.Method == "POST" {
		formFile, header, err := res.FormFile("in")
		if err != nil {
			log.Printf("create file no suc:%s", err)
			return
		}
		defer formFile.Close()
		save_file, err := os.Create("./" + header.Filename)
		if err != nil {
			return
		}
		defer save_file.Close()
		_, err = io.Copy(save_file, formFile)
		if err != nil {
			log.Printf("write file no suc:%s", err)
			return
		}

		// io.WriteString(w, "test"+time.Now().Format("2002-02-01"))
		fmt.Println("test" + time.Now().Format("2002-02-01"))
	}
}

func main() {
	http.HandleFunc("/", tfile)
	http.ListenAndServe(":8080", nil)

}
