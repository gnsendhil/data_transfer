package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Start client...")

	conn, _ := net.Dial("tcp", "127.0.0.1:9000")
	buf := make([]byte, 16)
	for {
		//message, _ := bufio.NewReader(conn).ReadString('\n')
		conn.Read(buf)
		fmt.Print("Received :", buf)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")
		if text == "end\n" {
			break
		}
	}
}
