package cells

import (
	"github.com/freecellsolver/cards"
	"slices"
)

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

func (fCell FieldCell) GetEndSeq() []cards.Card {
	if len(fCell.CardStack) == 0 {
		return []cards.Card{}
	}

	seq := make([]cards.Card, 0, len(fCell.CardStack))
	reversed := slices.Clone(fCell.CardStack)
	slices.Reverse(reversed)
	for _, card := range reversed {
		if len(seq) == 0 {
			seq = append(seq, card)
			continue
		}

		currentSeqRoot := seq[len(seq)-1]
		if currentSeqRoot.IsBlack() == card.IsBlack() {
			// If same color, end sequence
			break
		}

		if currentSeqRoot.GetNumCode()+1 != card.GetNumCode() {
			// If discontinued number, end sequence
			break
		}
		seq = append(seq, card)
	}

	slices.Reverse(seq)
	return seq
}
