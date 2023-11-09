package main

import (
	"io"
	"net/http"
)

const myfrom = `<html><body>
<form action="#" method="post" name="bar">
	<input type="text" name="zh" />
	<input type="text" name="mm" />
	<input type="submit" value="submit"/>
</form>
</body></html>`

func mywb(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch request.Method {
	case "GET":
		io.WriteString(w, myfrom)
	case "POST":
		zh := request.FormValue("zh")
		mm := request.FormValue("mm")
		if zh == "123" && mm == "123" {
			io.WriteString(w, "suc")

		} else {
			// io.WriteString(w, "fail")
			// time.Sleep(2 * time.Second)
			http.Redirect(w, request, "http://www.bilibili.com", http.StatusFound)
		}

	}
}

func main() {
	http.HandleFunc("/test1", mywb)
	http.ListenAndServe(":8080", nil)
}
