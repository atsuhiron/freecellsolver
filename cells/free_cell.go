package cells

import (
	"github.com/freecellsolver/cards"
	"slices"
)

type FreeCell struct {
	CardStack []cards.Card
}

func (fCell FreeCell) CanPlace(card cards.Card) bool {
	return len(fCell.CardStack) == 0
}

func (fCell FreeCell) GetEndSeq() []cards.Card {
	if len(fCell.CardStack) == 0 {
		return make([]cards.Card, 0)
	}
	return slices.Clone(fCell.CardStack)
}

func (fCell FreeCell) Clone() FreeCell {
	cloneStack := make([]cards.Card, len(fCell.CardStack))
	copy(cloneStack, fCell.CardStack)
	return FreeCell{cloneStack}
}
