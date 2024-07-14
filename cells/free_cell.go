package cells

import "github.com/freecellsolver/cards"

type FreeCell struct {
	CardStack []cards.Card
}

func (fCell FreeCell) CanPlace(card cards.Card) bool {
	return len(fCell.CardStack) == 0
}
