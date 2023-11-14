package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	rfe, _ := os.ReadFile("test.txt")
	fmt.Fprintf(w, string(rfe))
}
func main() {

	err := http.ListenAndServe(":8080", http.HandlerFunc(HelloServer))
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
