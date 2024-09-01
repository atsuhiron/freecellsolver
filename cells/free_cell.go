package cells

import (
	"fmt"
	"github.com/freecellsolver/cards"
	"slices"
)

type FreeCell struct {
	CardStack []cards.Card
}

func (fCell *FreeCell) CanPlace(_ cards.Card) bool {
	return len(fCell.CardStack) == 0
}

func (fCell *FreeCell) GetEndSeq() []cards.Card {
	if len(fCell.CardStack) == 0 {
		return make([]cards.Card, 0)
	}
	return slices.Clone(fCell.CardStack)
}

func (fCell *FreeCell) RemoveEndSeq(removeNum int) error {
	if len(fCell.CardStack) < removeNum {
		return fmt.Errorf("removeNum (%v) must be equal or samller than CardStack size (%v)", removeNum, len(fCell.CardStack))
	}
	fCell.CardStack = fCell.CardStack[:len(fCell.CardStack)-removeNum]
	return nil
}

func (fCell *FreeCell) Clone() FreeCell {
	cloneStack := make([]cards.Card, len(fCell.CardStack))
	copy(cloneStack, fCell.CardStack)
	return FreeCell{cloneStack}
}
