package cards

import (
	"fmt"
	"github.com/freecellsolver/consts"
	"strconv"
)

type ReadableCard struct {
	Suit string
	Num  string
}

func (rCard *ReadableCard) ToCard() (Card, error) {
	suitCode := uint8(0)
	switch rCard.Suit {
	case "♠":
		suitCode = 0
	case "♥":
		suitCode = 1
	case "♣":
		suitCode = 2
	case "♦":
		suitCode = 3
	case "S":
		suitCode = 0
	case "H":
		suitCode = 1
	case "C":
		suitCode = 2
	case "D":
		suitCode = 3
	default:
		return Card{uint8(255)}, fmt.Errorf("failed to convert suit: %s", rCard.Suit)
	}

	numCode, err := strconv.Atoi(rCard.Num)
	if err != nil {
		switch rCard.Num {
		case "A":
			numCode = 1
		case "J":
			numCode = 11
		case "Q":
			numCode = 12
		case "K":
			numCode = 13
		case "T":
			numCode = 10

		default:
			return Card{uint8(255)}, fmt.Errorf("failed to convert number: %s", rCard.Num)
		}
	} else {
		if numCode < 0 || numCode > 13 {
			return Card{uint8(255)}, fmt.Errorf("invalid number: %s", rCard.Num)
		}
	}

	return Card{suitCode<<consts.SShift + uint8(numCode)}, nil
}
