package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Status struct {
	Text string
}
type User struct {
	XMLName xml.Name
	Status  Status
}

func main() {
	res, err := http.Get("http://twitter.com/users/Googland.xml")
	checkErr(err)
	// data, err := ioutil.ReadAll(res.Body)
	checkErr(err)
	// fmt.Printf("Ger: %q", string(data))
	user := User{xml.Name{"", "user"}, Status{""}}
	boy, _ := io.ReadAll(res.Body)
	c := xml.Unmarshal(boy, &user)

	fmt.Printf("status:%s", c)
}
func checkErr(err error) {
	if err != nil {
		log.Fatalf("Get:%v", err)
	}
}
