package cells

import "github.com/freecellsolver/cards"

type FreeCell struct {
	cardStack []cards.Card
}

func (fCell FreeCell) CanPlace(card cards.Card) bool {
	return len(fCell.cardStack) == 0
}
