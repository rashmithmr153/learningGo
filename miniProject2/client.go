package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func copy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		fmt.Print(err)
		return
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Print(err)
		return
	}

	defer conn.Close()
	go copy(os.Stdout, conn)
	fmt.Println("Read or write from files present in server")
	fmt.Println("command format [opertion] [filename] [string]")
	fmt.Println("[opertion]: r(for read), w(for write)")
	reader := bufio.NewReader(os.Stdin)

	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Input error:", err)
			return
		}

		_, err = conn.Write([]byte(text))
		if err != nil {
			fmt.Println("Connection closed:", err)
			return
		}
	}
}
