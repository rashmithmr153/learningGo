package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func Copy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		return
	}
}

func handConc(conec net.Conn) {
	defer conec.Close()
	fmt.Println("Server started")
	for {
		_, err := io.WriteString(conec, "----------\n")
		if err != nil {
			return
		}
		Copy(os.Stdout, conec)
		time.Sleep(time.Second)
	}
}

func main() {
	fmt.Println("Server starting at port 8080")
	listen, err := net.Listen("tcp", "loalhost:8080")
	if err != nil {
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		go handConc(conn)
	}
}
