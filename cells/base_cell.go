package cells

import (
	"github.com/freecellsolver/cards"
)

type BaseCell interface {
	CanPlace(card cards.Card) bool
	GetEndSeq() []cards.Card
}
