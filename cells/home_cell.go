package cells

import "github.com/freecellsolver/cards"

type HomeCell struct {
	CardStack []cards.Card
	SuitCode  uint8 // right shifted
}

func (hCell HomeCell) CanPlace(card cards.Card) bool {
	if card.GetSuitCode() != hCell.SuitCode {
		return false
	}

	return card.Code == hCell.CardStack[len(hCell.CardStack)-1].Code+1
}
