package main

import (
	"fmt"
	"math/rand"
)

type Player interface {
	KickBall()
}

type FootballPlayer struct {
	stamina int
	power   int
}

type CR7 struct {
	stamina int
	power   int
	SUI     int
}

type Messi struct {
	FootballPlayer
	WCbonus int
}

func (f CR7) KickBall() {
	shot := f.stamina + f.power*f.SUI
	fmt.Println("CR7 is kicking the ball", shot)
}

func (m Messi) KickBall() {
	shot := (m.stamina + m.power) * m.WCbonus
	fmt.Println("Messi is kicking the ball", shot)
}

func (f FootballPlayer) KickBall() {
	shot := f.stamina + f.power
	fmt.Println("Random is kicking the ball", shot)
}

func main() {
	team := make([]Player, 18)
	for i := 0; i < len(team)-2; i++ {
		team[i] = FootballPlayer{
			stamina: rand.Intn(10),
			power:   rand.Intn(10),
		}
	}
	team[len(team)-1] = CR7{
		stamina: 10,
		power:   10,
		SUI:     10,
	}
	team[len(team)-2] = Messi{
		FootballPlayer: FootballPlayer{
			stamina: 10,
			power:   8,
		},
		WCbonus: 10,
	}
	for i := 0; i < len(team); i++ {
		team[i].KickBall()
	}
}
