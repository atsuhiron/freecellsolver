package cells

import (
	"github.com/freecellsolver/cards"
)

type BaseCell interface {
	CanPlace(card cards.Card) bool
	GetEndSeq(onlyLast bool) []cards.Card
	RemoveEndSeq(removeNum int) error
}
