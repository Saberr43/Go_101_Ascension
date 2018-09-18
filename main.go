package main

import (
	"fmt"
	"pokerdeck/dealers"
	"pokerdeck/players"
)

func main() {
	player1 := player.Player{Name: "Player 1"}
	player2 := player.Player{Name: "Player 2"}

	dealer := dealer.Dealer{}

	var err error
	dealer.DealerDeck, err = dealer.Deal(dealer.DealerDeck, &player1)
	if err != nil {
		panic(err)
	}

	dealer.DealerDeck.PrintDeck()

	fmt.Printf("%+v", player1.Cards)

	dealer.DealerDeck, err = dealer.Deal(dealer.DealerDeck, &player2)
	if err != nil {
		panic(err)
	}

	println("-----------------------------------------")

	dealer.DealerDeck.PrintDeck()

	fmt.Printf("%+v", player2.Cards)
}
