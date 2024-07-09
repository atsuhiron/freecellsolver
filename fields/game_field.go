package fields

import (
	"fmt"
	"slices"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/freecellsolver/cards"
	"github.com/freecellsolver/cells"
)

var suits = []uint8{0, 1, 2, 3}

type GameField struct {
	Homes  map[uint8]cells.HomeCell
	Frees  [4]cells.FreeCell
	Fields [8]cells.FieldCell
}

func CreateGameField(fie [8][]cards.Card, fre [4][]cards.Card, hom map[uint8][]cards.Card) (GameField, error) {
	allCards := slices.Concat(
		fie[0], fie[1], fie[2], fie[3], fie[4], fie[5], fie[6], fie[7],
		fre[0], fre[1], fre[2], fre[3],
		hom[uint8(0)], hom[uint8(1)], hom[uint8(2)], hom[uint8(3)],
	)

	if len(allCards) != 52 {
		return GameField{}, fmt.Errorf("the number of cards is not 52: %d", len(allCards))
	}

	if !checkUniqCards(allCards) {
		return GameField{}, fmt.Errorf("cards are not unique")
	}

	if !checkHomeCell(hom) {
		return GameField{}, fmt.Errorf("cards in home cell are invalid")
	}

	if !checkFreeCell(fre) {
		return GameField{}, fmt.Errorf("cards in free cell are invalid")
	}

	return GameField{}, nil
}

func checkUniqCards(cards []cards.Card) bool {
	cardSet := mapset.NewSet[uint8]()

	for _, card := range cards {
		cardSet.Add(card.Code)
	}

	return len(cards) == cardSet.Cardinality()
}

func checkHomeCell(hom map[uint8][]cards.Card) bool {
	suitCodeSet := mapset.NewSet[uint8](suits...)

	for kSuitCode, vCell := range hom {
		if !suitCodeSet.Contains(kSuitCode) {
			return false
		}

		inCellSuitSet := mapset.NewSet[uint8]()
		for _, card := range vCell {
			inCellSuitSet.Add(card.GetSuitCode())
		}

		if inCellSuitSet.Cardinality() > 1 {
			return false
		} else if inCellSuitSet.Cardinality() == 1 {
			if vCell[0].GetSuitCode() != kSuitCode {
				return false
			}
		}
	}

	return true
}

func checkFreeCell(fre [4][]cards.Card) bool {
	for _, cell := range fre {
		if len(cell) > 1 {
			return false
		}
	}
	return true
}
