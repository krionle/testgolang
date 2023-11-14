package main

import (
	"encoding/json"
	"filestore-server/meta"
	"filestore-server/util"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func tfile(w http.ResponseWriter, res *http.Request) {
	if res.Method == "GET" {
		data, err := io.ReadFile("tfile.html")
		if err != nil {
			io.WriteString(data, "html 未找到")
			return
		}
	} else if res.Method == "POST" {
		file, head, err := res.FormFile("file")
		if err != nil {
			fmt.Print("no file")
			return
		}
		defer file.Close()
		fileMeta := meta.fileMeta{
			FileName: head.Filename,
			Location: "./" + head.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		newFile, err := os.Create(fileMeta.Location)
		if err != nil {
			fmt.Printf("Filed to create")
			return
		}
		defer newFile.Close()
		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Print("File to save")
			return
		}
		newFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		fmt.Print(fileMeta.FileSha1)
		meta.UpdateFileMeta(fileMeta)
		http.Redirect(w, res, "./", http.StatusFound)
	}
}

func main() {
	http.HandleFunc("/", tfile)
	http.ListenAndServe(":8080", nil)

}
