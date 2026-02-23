package main

import (
	"fmt"
	"net"
	"sync"
)

func checkWin(game *Game) {
	p1 := &game.Players[0]
	p2 := &game.Players[1]
	defer p1.conn.Close()
	defer p2.conn.Close()
	switch game.Winner {
	case p1:
		p1.conn.Write([]byte("Yay...You Win!\n"))
		p2.conn.Write([]byte("You Lose!.Better luck next time.\n"))
		return
	case p2:
		p2.conn.Write([]byte("Yay...You Win!\n"))
		p1.conn.Write([]byte("You Lose!.Better luck next time.\n"))
		return
	default:
		p1.conn.Write([]byte("Draw,as both took same number of guesses\n"))
		p2.conn.Write([]byte("Draw,as both took same number of guesses\n"))
		return
	}
}

func main() {
	var gameDetails Game
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
		gameDetails.Players = append(gameDetails.Players, playerDetails)
		conn.Write([]byte("___________________________________________________\n"))
		conn.Write([]byte("*************Welcome to bull-cow game**************\n"))
		conn.Write([]byte("___________________________________________________\n"))
	}
	player1 := &gameDetails.Players[0]
	player2 := &gameDetails.Players[1]
	gameDetails.secretNo = SectNumGenrator()
	fmt.Println(gameDetails.secretNo)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		handlePlayer(&gameDetails, player1)
	}()

	go func() {
		defer wg.Done()
		handlePlayer(&gameDetails, player2)
	}()

	wg.Wait()
	checkWin(&gameDetails)
}
