package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("http.ser is staring")
	listener, err := net.Listen("tcp", "localhost:80")
	if err != nil {
		print("err")
	}
	for {
		coon, err := listener.Accept()
		if err != nil {
			fmt.Print("Err")
			return
		}
		go dohtser(coon)
	}
}

func dohtser(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Print("err read", err.Error())
			return
		}
		fmt.Printf("recievied date:%v", string(buf[:len]))
	}
}
