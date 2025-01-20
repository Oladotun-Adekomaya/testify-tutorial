package main

import (
	"fmt"
	"sync"
	// "math/rand"
	// "sync"
)

type Player interface {
	KickBall()
}

type FootballPlayer struct {
	stamina int
	power   int
}

func (f FootballPlayer) KickBall() {
	shot := f.stamina + f.power
	fmt.Println("i'm kicking the ball", shot)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	wg.Done()

	wg.Wait()
}
