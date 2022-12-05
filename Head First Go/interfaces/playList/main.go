package main

import (
	"tapes/cassette"
)

type Player interface {
	Play(string)
	Stop()
}

func main() {
	var player Player = cassette.TapePlayer{}
	mixTape := []string{"Viktor and his demons", "Mosca Cojonera", "Chop suey"}
	playList(player, mixTape)
	player = cassette.TapeRecorder{}
	playList(player, mixTape)
	TryOut(cassette.TapeRecorder{})

}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

// Type assertion

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	recorder := player.(cassette.TapeRecorder)
	recorder.Record()
}
