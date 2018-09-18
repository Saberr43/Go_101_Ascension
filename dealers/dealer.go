package dealer

import (
	"errors"
	"pokerdeck/decks"
	"pokerdeck/players"
)

//Dealer deals cards to players
type Dealer struct {
	DealerDeck deck.Deck //DealerDeck is the dealers instance of the deck
}

//Deal gives players cards from deck object
func (dealer Dealer) Deal(d deck.Deck, p *player.Player) (deck.Deck, error) {
	if len(d) == 0 {
		d = deck.NewDeck()
		d.Shuffle()
	}

	if !p.IsCardsPresent() {
		var err error
		p.Cards, err = take2OffTopOfStringSlice(&d)

		if err != nil {
			panic(err)
		}

		return d, nil
	}

	return d, errors.New("player has cards present")
}

func take2OffTopOfStringSlice(d *deck.Deck) ([2]string, error) {
	var result [2]string

	if len(*d) < 2 {
		return result, errors.New("Not enough cards in deck")
	}

	result[0] = (*d)[0]
	result[1] = (*d)[1]

	*d = (*d)[2:]

	return result, nil
}
