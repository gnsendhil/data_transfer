package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

//commit2
func main() {
	fmt.Println("Start server...")

	ln, _ := net.Listen("tcp", ":9000")

	conn, _ := ln.Accept()
	fmt.Println("Accepted connection")

	f, err := os.Open("fileread_write.go")

	if err != nil {
		log.Fatal("no file exists")
	}

	defer f.Close()
	defer conn.Close()

	buf := make([]byte, 16)

	for {
		n, err := f.Read(buf)

		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF")
				break
			} else {
				log.Fatal("error while reading")
			}
		}

		fmt.Print(buf)
		fmt.Println(string(buf[0:n]))
		conn.Write(buf)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		//fmt.Printf("Message Receied %s\n", string(message))
		if message == "end\n" {
			break
		}
	}
}
//commit3
//commit4
