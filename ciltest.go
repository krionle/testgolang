package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "loaclhost:80")
	if err != nil {
		print("目标繁忙\n")
		return
	}
	inreader := bufio.NewReader(os.Stdin)
	fmt.Println("1,what is na")
	cliname, _ := inreader.ReadString('\n')
	triClient := strings.Trim(cliname, "\r\n")
	for {
		fmt.Println("Q is quit")
		input, _ := inreader.ReadString('\n')
		triInput := strings.Trim(input, "\r\n")
		if triInput == "Q" {
			return
		}
		_, err := conn.Write([]byte(triClient + "say:" + triInput))
		println(err)
	}
}
