package player

//Player represents one card player
type Player struct {
	Name  string
	Cards [2]string
}

//IsCardsPresent checks if the player has cards
func (p Player) IsCardsPresent() bool {
	if len(p.Cards[0]) > 0 && len(p.Cards[1]) > 0 {
		return true
	}

	return false
}
