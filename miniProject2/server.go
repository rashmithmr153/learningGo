package main

import (
	"fmt"
	"net"
)

func handConc(conec net.Conn) {
	defer conec.Close()
	fmt.Println("Client connceted;",&conec)
	b := make([]byte, 1024)
	for {

		n, err := conec.Read(b)
		if err != nil {
			fmt.Print(err)
			return
		}
		msg := string(b[:n])
		fmt.Println("client:", msg)
		resp := "server repose:" + msg + "recived sucessfuly\n"
		conec.Write([]byte(resp))
		// time.Sleep(time.Second)
	}
}

func main() {
	fmt.Println("Server starting at port 8080")
	listen, err := net.Listen("tcp", "localhost:8080")
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
