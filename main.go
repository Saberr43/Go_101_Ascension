package main

import "pokerdeck/decks"

func main() {
	deck := deck.NewDeck()

	deck.Shuffle()

	deck.PrintDeck()
}
