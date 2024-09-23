package main

import (
	"fmt"

	"github.com/freecellsolver/cards"
	"github.com/freecellsolver/cells"
)

/*
type Card struct {
	Code uint8
}

func genCards() []Card {
	var cards [52]Card
	for i := 0; i < 52; i++ {
		num := uint8(i % 13)
		suit := uint8(i / 13)
		cards[i] = Card{suit*32 + num}
	}
	return cards[:]
}
*/

func main() {
	ccard := cards.Card{Code: uint8(42)}
	fmt.Printf("%v\n", ccard)
	rcard, err := ccard.ToReadableCard()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v\n", rcard)

	freeCells := make([]cells.FreeCell, 4)
	freeCells = append(freeCells, cells.FreeCell{})

	fmt.Printf("Can place card: %v", freeCells[0].CanPlace(cards.Card{Code: 37}))
}
