package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Println("Start client...")

	wf, err := os.Create("out.go")

	if err != nil {
		log.Fatal(err)
	}
	defer wf.Close()

	conn, _ := net.Dial("tcp", "127.0.0.1:9000")
	buf := make([]byte, 16)
	for {
		_, err := conn.Read(buf)

		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF")
				break
			} else {
				log.Fatal("error while reading")
			}
			break
		}

		fmt.Print("Received :", buf)

		//fmt.Printf("got %d bytes", n)
		str := string(buf[0:16])
		fmt.Println(str)

		_, err1 := wf.Write(buf)

		if err1 != nil {
			log.Fatal(err1)
		}

		//reader := bufio.NewReader(os.Stdin)
		//fmt.Print("Text to send: ")
		//text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, "ACK"+"\n")
	}
}
