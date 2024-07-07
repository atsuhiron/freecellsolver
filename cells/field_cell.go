package cells

import "github.com/freecellsolver/cards"

type FieldCell struct {
	CardStack []cards.Card
}

func (fCell FieldCell) CanPlace(card cards.Card) bool {
	if len(fCell.CardStack) == 0 {
		return true
	}

	if fCell.CardStack[len(fCell.CardStack)-1].IsBlack() == card.IsBlack() {
		// same color
		return false
	}
	return card.GetNumCode() == fCell.CardStack[len(fCell.CardStack)-1].GetNumCode()-uint8(1)
}
