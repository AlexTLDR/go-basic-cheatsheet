package main

import (
	"fmt"
	"math/rand"
)

type Player interface {
	KickBall() int
	Name() string
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

func (c CR7) KickBall() int {
	return c.stamina + c.power*c.SUI

}

func (c CR7) Name() string {
	return "CR7"
}

func (m Messi) KickBall() int {
	return (m.stamina + m.power) * m.WCbonus

}

func (m Messi) Name() string {
	return "Messi"
}

func (f FootballPlayer) KickBall() int {
	return f.stamina + f.power
}

func (f FootballPlayer) Name() string {
	return "Random dude"
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
		fmt.Printf("%s is kicking the ball %d\n", team[i].Name(), team[i].KickBall())
	}
}
