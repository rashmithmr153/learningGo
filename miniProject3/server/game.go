package main

import (
	"math/rand"
	"net"
	"strings"
	"strconv"
	"sync"
)
type Game struct{
	secretNo string
	Players [] Player
	Winner *Player
	isGameover bool
	Mutex sync.Mutex
}

type Player struct{
	conn net.Conn
	Id int
	guessCount int
	isFinished bool
}


func SectNumGenrator() string{
	var secret string
	for len(secret)<4{
		num:=rand.Intn(10)
		strNum := strconv.Itoa(num)
		if !strings.Contains(secret,strNum){
			secret+=strNum
		}
	}
	return secret
}

func handlePlayer(game *Game,player *Player){

}