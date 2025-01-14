package main

import (
	"fmt"
	"math/rand"
)

type FootballPlayer struct {
	stamina int
	power   int
}

func (f FootballPlayer) KickBall() {
	shot := f.stamina + f.power
	fmt.Println("i'm kicking the ball", shot)
}

func main() {
	team := make([]FootballPlayer, 11)

	team[10] = FootballPlayer{
		stamina: 10,
		power:   7,
	}
	for i := 0; i < len(team); i++ {
		team[i] = FootballPlayer{
			stamina: rand.Intn(10),
			power:   rand.Intn(10),
		}
	}
	for i := 0; i < len(team); i++ {
		team[i].KickBall()
	}
}
