package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Start server...")

	ln, _ := net.Listen("tcp", ":9000")

	conn, _ := ln.Accept()

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Printf("Message Receied %s\n", string(message))
		if message == "end\n" {
			break
		}
	}
}
