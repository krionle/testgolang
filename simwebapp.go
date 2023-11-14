package main

import (
	"io"
	"net/http"
)

const form = `
    <html><body>
        <form action="#" method="post" name="bar">
            <input type="text" name="in" />
            <input type="submit" value="submit"/>
        </form>
    </body></html>
`

func swap(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "<h1>hello</h1>")
}
func FormServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("content-type", "text/html")
	switch request.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		io.WriteString(w, request.FormValue("in"))
	}

}
func main() {
	http.HandleFunc("/test1", swap)
	http.HandleFunc("/test2", FormServer)
	http.ListenAndServe(":8080", nil)
}
