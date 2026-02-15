package main

import (
	// "errors"
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func fileOpertions(command string, conn net.Conn) {
	// buff:=make([]byte,1024)
	splitCmd := strings.Fields(command)
	if len(splitCmd) < 2 {
		b := []byte("Invalid command refer command format\n")
		conn.Write(b)
		return
	}
	opertion := splitCmd[0]
	filename := splitCmd[1]
	if opertion == "w" {
		if len(splitCmd) < 3 {
			b := []byte("Invalid command refer command format\n")
			conn.Write(b)
			return
		}
	}
	str := strings.Join(splitCmd[2:], " ")

	switch opertion {
	case "r":
		b, err := os.ReadFile(filename)
		if err != nil {
			resp := []byte(err.Error() + "\n")
			conn.Write(resp)
			return
		}
		conn.Write(b)
		return
	case "w":
		file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			b := []byte(err.Error() + "\n")
			conn.Write(b)
			return
		}
		defer file.Close()
		file.WriteString(str + "\n")
		resp := "wrie opertinon done\n"
		conn.Write([]byte(resp))
		return
	default:
		resp := "'" + string(opertion) + "' " + "Invalid file opertions\n"
		conn.Write([]byte(resp))
		return
	}
}

func handConc(conec net.Conn) {
	defer conec.Close()
	fmt.Println("Client connceted;", conec.RemoteAddr())
	// b := make([]byte, 1024)
	reader := bufio.NewReader(conec)
	for {
		// n, err := conec.Read(b)
		command, err := reader.ReadString('\n')
		command = strings.TrimSpace(command)
		if err != nil {
			fmt.Print(err)
			conec.Close()
			return
		}
		fileOpertions(command, conec)
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
