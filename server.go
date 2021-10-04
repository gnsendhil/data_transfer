package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

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

	reader := bufio.NewReader(f)

	buf := make([]byte, 16)

	for {
		_, err := reader.Read(buf)

		if err != nil {
			if err != io.EOF {
				log.Fatal("error while reading")
			}
			break
		}
		//conn.Write(buf)
		fmt.Fprint(conn, buf)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Printf("Message Receied %s\n", string(message))
		if message == "end\n" {
			break
		}
	}
}
