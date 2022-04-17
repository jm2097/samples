package main

import "fmt"

type Player struct {
	Score int
}

func main() {
	player := Player{
		Score: 0,
	}

	fmt.Println("Before", player.Score)
	player.addScore(1)
	fmt.Println("After", player.Score)

	player.resetScore()
	fmt.Println("Reset", player.Score)
}

func (p *Player) addScore(score int) {
	p.Score += score
}

func (p *Player) resetScore() {
	p.Score = 0
}
