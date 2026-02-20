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
func calcBullCows(secreatNo,guess string) (bull,cow int){
	for i:=0;i<4;i++{
		if secreatNo[i]==guess[i]{
			bull+=1
		}else if strings.Contains(secreatNo,guess[i:i+1]){
			cow+=1
		}
	}
	return
}

func handlePlayer(game *Game, player *Player) {
	conec := player.conn
	reader := bufio.NewReader(conec)
	for player.guessCount<5 {
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
		if game.isGameover {
			game.lock.Unlock()
			return
		}
		bulls,cows:=calcBullCows(game.secretNo,guess)
		if bulls==4{
			game.isGameover=true
			game.Winner=player
			game.lock.Unlock()
			resp:="Yayy... your the winner\n"
			return
		}
		resp:="No of bulls-->"+strconv.Itoa(bulls)+", cows-->"+strconv.Itoa(cows)+"\n"
		player.guessCount+=1
		game.lock.Unlock()
		conec.Write([]byte(resp))
	}

}
