package cards

import (
	"fmt"
	"strconv"
)

type Card struct {
	Code uint8
}

func (card Card) GetSuitCode() uint8 {
	return card.Code >> 5
}

func (card Card) IsBlack() bool {
	return card.GetSuitCode()%2 == 0
}

func (card Card) ToReadableCard() (ReadableCard, error) {
	suitCode := card.GetSuitCode()
	suit := ""
	switch suitCode {
	case 0:
		suit = "♠"
	case 1:
		suit = "♥"
	case 2:
		suit = "♣"
	case 3:
		suit = "♦"
	default:
		return ReadableCard{}, fmt.Errorf("failed to convert suit: %d", suitCode)
	}

	numCode := card.Code - suitCode<<5
	num := ""
	switch numCode {
	case 1:
		num = "A"
	case 11:
		num = "J"
	case 12:
		num = "Q"
	case 13:
		num = "K"
	default:
		if (2 <= numCode) && (numCode <= 10) {
			num = strconv.Itoa(int(numCode))
		} else {
			return ReadableCard{}, fmt.Errorf("failed to convert number: %d", numCode)
		}
	}

	return ReadableCard{suit, num}, nil
}

func CreateCard(str string) (Card, error) {
	if len(str) < 2 {
		return Card{uint8(255)}, fmt.Errorf("cannot parse to Card: %s", str)
	}

	rc := ReadableCard{Suit: string(str[0]), Num: str[1:]}
	c, err := rc.ToCard()
	if err != nil {
		return Card{uint8(255)}, fmt.Errorf("cannot parse to Card: %s", str)
	}
	return c, nil
}
