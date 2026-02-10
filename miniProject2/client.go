package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func copy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		return
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		return
	}

	defer conn.Close()
	copy(os.Stdout, conn)
	for {
		var sndStr string
		fmt.Scan("Enter some string: ", &sndStr)
		io.WriteString(conn, sndStr)
	}
}
