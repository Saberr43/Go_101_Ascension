package main

import (
	"pokerdeck/players"
)

func main() {
	player := player.Player{Name: "Player 1", Cards: [2]string{"A of Spades", "A of Diamonds"}}
	print(player.IsCardsPresent())
}
