package main

import (
	"fmt"
	"net"
)

func main() {
	var gameDetails Game
	fmt.Println("___________________________________________________")
	fmt.Println("*************Welcomr to bull-cow game**************")
	fmt.Println("___________________________________________________")
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
	}
	for len(gameDetails.Players) < 2 {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error: ", err.Error())
		}
		var playerDetails Player
		playerDetails.conn = conn
		playerDetails.guessCount = 0
		playerDetails.Id = len(gameDetails.Players) + 1
		playerDetails.isFinished = false

		gameDetails.Players = append(gameDetails.Players, playerDetails)
	}
	gameDetails.secretNo = SectNumGenrator()
	fmt.Println(gameDetails.secretNo)
	go handlePlayer(&gameDetails, &gameDetails.Players[0])
	go handlePlayer(&gameDetails, &gameDetails.Players[1])
	select {}
}
