package main

import (
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
	for {
		var sndStr string
		fmt.Println("Enter some string:")
		fmt.Scan(&sndStr)
		conn.Write([]byte(sndStr + "\n"))
	}
}
