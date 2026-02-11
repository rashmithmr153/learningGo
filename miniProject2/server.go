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
	opertion := splitCmd[0]
	filename := splitCmd[1]
	var str string
	if opertion == "w" {
		str = splitCmd[2]
	} else {
		str = ""
	}

	switch opertion {
	case "r":
		b, err := os.ReadFile(filename)
		if err != nil {
			fmt.Print("Error: ", err)
			return
		}
		conn.Write(b)
		return
	case "w":
		file, err := os.OpenFile(filename, os.O_APPEND, 0666)
		if err != nil {
			fmt.Print(err)
		}
		file.WriteString(str + "\n")
		file.Close()
		resp := "wrie opertinon done\n"
		conn.Write([]byte(resp))
		return
	default:
		resp := "'"+string(opertion)+"' "+ "Invalid file opertions\n"
		conn.Write([]byte(resp))
		return
	}
}

func handConc(conec net.Conn) {
	defer conec.Close()
	fmt.Println("Client connceted;", &conec)
	// b := make([]byte, 1024)
	reader := bufio.NewReader(conec)
	for {
		// n, err := conec.Read(b)
		command, err := reader.ReadString('\n')
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
