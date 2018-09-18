package deck

import (
	"fmt"
	"math/rand"
	"time"
)

//Deck is an array of strings
type Deck []string

//NewDeck returns a new deck
func NewDeck() Deck {
	resultDeck := Deck{}
	cardValues := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
	cardSuits := []string{"Diamond", "Club", "Heart", "Spade"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			resultDeck = append(resultDeck, value+" of "+suit)
		}
	}
	return resultDeck
}

//PrintDeck prints the cards in a deck
func (d Deck) PrintDeck() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//Shuffle shuffles a deck
func (d Deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPos := r.Intn(len(d) - 1)

		d[i], d[newPos] = d[newPos], d[i]
	}
}
