package main

import (
	"bufio"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"sync"
)

type Game struct {
	secretNo   string
	Players    []Player
	Winner     *Player
	isGameover bool
	lock       sync.Mutex
}

type Player struct {
	conn       net.Conn
	Id         int
	guessCount int
	isFinished bool
}

func SectNumGenrator() string {
	var secret string
	for len(secret) < 4 {
		num := rand.Intn(10)
		strNum := strconv.Itoa(num)
		if !strings.Contains(secret, strNum) {
			secret += strNum
		}
	}
	return secret
}

func validGuess(guess string) bool {
	if len(guess) != 4 {
		return false
	}
	seen := make(map[byte]bool)

	for i := range guess {
		if guess[i] < '0' || guess[i] > '9' {
			return false
		}

		if seen[guess[i]] {
			return false
		}

		seen[guess[i]] = true
	}
	return true
}

func handlePlayer(game *Game, player *Player) {
	conec := player.conn
	reader := bufio.NewReader(conec)
	for {
		guess, err := reader.ReadString('\n')
		if err != nil {
			conec.Write([]byte(err.Error()))
			return
		}
		guess = strings.TrimSpace(guess)
		if !validGuess(guess) {
			conec.Write([]byte("Enter valid guess"))
			continue
		}
		game.lock.Lock()
		//call fun to check guess and retrn bulls and cows

	}

}
