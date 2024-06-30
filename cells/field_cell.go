package cells

import "github.com/freecellsolver/cards"

type FieldCell struct {
	cardStack []cards.Card
}

func (fCell FieldCell) CanPlace(card cards.Card) bool {
	if len(fCell.cardStack) == 0 {
		return true
	}

	return fCell.cardStack[len(fCell.cardStack)-1].IsBlack() != card.IsBlack() // xor
}
