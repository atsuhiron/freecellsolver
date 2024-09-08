package cells

import (
	"fmt"
	"github.com/freecellsolver/cards"
)

type HomeCell struct {
	CardStack []cards.Card
	SuitCode  uint8 // right shifted
}

func (hCell *HomeCell) CanPlace(card cards.Card) bool {
	if card.GetSuitCode() != hCell.SuitCode {
		return false
	}
	if len(hCell.CardStack) == 0 {
		return card.Code%32 == 1
	}
	return card.Code == hCell.CardStack[len(hCell.CardStack)-1].Code+1
}

func (hCell *HomeCell) GetEndSeq() []cards.Card {
	return make([]cards.Card, 0)
}

func (hCell *HomeCell) RemoveEndSeq(removeNum int) error {
	if len(hCell.CardStack) < removeNum {
		return fmt.Errorf("removeNum (%v) must be equal or samller than CardStack size (%v)", removeNum, len(hCell.CardStack))
	}
	hCell.CardStack = hCell.CardStack[:len(hCell.CardStack)-removeNum]
	return nil
}

func (hCell *HomeCell) Place(seq *[]cards.Card) {
	hCell.CardStack = append(hCell.CardStack, *seq...)
}

func (hCell *HomeCell) Clone() *HomeCell {
	cloneStack := make([]cards.Card, len(hCell.CardStack))
	copy(cloneStack, hCell.CardStack)
	return &HomeCell{cloneStack, hCell.SuitCode}
}
