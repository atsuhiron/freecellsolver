package main

import (
	"fmt"
	"slices"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/freecellsolver/cards"
	"github.com/freecellsolver/cells"
)

type GameField struct {
	Homes  map[uint8]cells.HomeCell
	Frees  [4]cells.FreeCell
	Fields [8]cells.FieldCell
}

func CreateGameField(fie [8][]cards.Card, fre [4][]cards.Card, hom [4][]cards.Card) (GameField, error) {
	allCards := slices.Concat(
		fie[0], fie[1], fie[2], fie[3], fie[4], fie[5], fie[6], fie[7],
		fre[0], fre[1], fre[2], fre[3],
		hom[0], hom[1], hom[2], hom[3],
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

func checkHomeCell(hom [4][]cards.Card) bool {
	overCellSuitSlice := make([]uint8, 4)

	for _, cell := range hom {
		inCellSuitSet := mapset.NewSet[uint8]()
		for _, card := range cell {
			inCellSuitSet.Add(card.GetSuitCode())
		}

		if inCellSuitSet.Cardinality() > 1 {
			return false
		} else if inCellSuitSet.Cardinality() == 1 {
			overCellSuitSlice = append(overCellSuitSlice, cell[0].GetSuitCode())
		}
	}

	overCellSuitSet := mapset.NewSet[uint8]()
	for _, s := range overCellSuitSlice {
		overCellSuitSet.Add(s)
	}

	return overCellSuitSet.Cardinality() == len(overCellSuitSlice)
}

func checkFreeCell(fre [4][]cards.Card) bool {
	for _, cell := range fre {
		if len(cell) > 1 {
			return false
		}
	}
	return true
}
