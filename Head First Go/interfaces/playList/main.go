package main

import (
	"tapes/cassette"
)

func main() {
	player := cassette.TapePlayer{}
	mixTape := []string{"Viktor and his demons", "Mosca Cojonera", "Chop suey"}
	playList(player, mixTape)
}

func playList(device cassette.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}
