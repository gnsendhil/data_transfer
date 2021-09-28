package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	f, err := os.Open("server.go")

	if err != nil {
		log.Fatal("no file exists")
	}

	defer f.Close()

	wf, err := os.Create("out.go")

	if err != nil {
		log.Fatal(err)
	}
	defer wf.Close()

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

		_, err1 := wf.Write(buf)

		if err1 != nil {
			log.Fatal(err1)
		}

		fmt.Printf("%s", buf)
	}
}
