package main

import (
	"fmt"
	"math/rand"
)

type FootBallPlayer struct {
	stamina int
	power   int
}

func (f FootBallPlayer) KickBall() {
	shot := f.stamina + f.power
	fmt.Println("I'm kicking the ball", shot)
}

func implement() {
	team := make([]FootBallPlayer, 11)
	for i := 0; i < len(team); i++ {
		team[i] = FootBallPlayer{
			stamina: rand.Intn(10),
			power:   rand.Intn(10),
		}
	}

	for _, footBallPlayer := range team {
		footBallPlayer.KickBall()
	}
}
