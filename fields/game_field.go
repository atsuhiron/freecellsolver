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

func CreateGameField(fie [][]cards.Card, fre [][]cards.Card, hom map[uint8][]cards.Card) (GameField, error) {
	// validate
	if !checkHomeCell(hom) {
		return GameField{}, fmt.Errorf("cards in home cell are invalid")
	}

	if !checkFreeCell(fre) {
		return GameField{}, fmt.Errorf("cards in free cell are invalid")
	}

	if len(fie) != 8 {
		return GameField{}, fmt.Errorf("length of field cell are invalid")
	}
	if len(fre) != 4 {
		return GameField{}, fmt.Errorf("length of free cell are invalid")
	}
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

	// convert
	//homes := make(map[uint8]cells.HomeCell)
	return GameField{}, nil
}

func CreateGameFieldFromReadable(
	fieR [][]cards.ReadableCard,
	freR [][]cards.ReadableCard,
	homR map[uint8][]cards.ReadableCard) (GameField, error) {
	fie := make([][]cards.Card, 0, len(fieR))
	for i, cellR := range fieR {
		cell := make([]cards.Card, 0, len(cellR))
		for _, cardR := range cellR {
			card, err := cardR.ToCard()
			if err != nil {
				return GameField{}, err
			}
			cell = append(cell, card)
		}
		fie[i] = cell
	}

	fre := make([][]cards.Card, 0, len(freR))
	for i, cellR := range freR {
		cell := make([]cards.Card, 0, len(cellR))
		for _, cardR := range cellR {
			card, err := cardR.ToCard()
			if err != nil {
				return GameField{}, err
			}
			cell = append(cell, card)
		}
		fre[i] = cell
	}

	hom := make(map[uint8][]cards.Card)
	for _, suitCode := range suits {
		cell := make([]cards.Card, 0, len(homR[suitCode]))
		for _, cardR := range homR[suitCode] {
			card, err := cardR.ToCard()
			if err != nil {
				return GameField{}, err
			}
			cell = append(cell, card)
		}
		hom[suitCode] = cell
	}

	return CreateGameField(fie, fre, hom)
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

func checkFreeCell(fre [][]cards.Card) bool {
	for _, cell := range fre {
		if len(cell) > 1 {
			return false
		}
	}
	return true
}
