package cells

import "github.com/freecellsolver/cards"

func EqualStack(stack1, stack2 *[]cards.Card) bool {
	if len(*stack1) != len(*stack2) {
		return false
	}
	for i := range *stack1 {
		if (*stack1)[i].Code != (*stack2)[i].Code {
			return false
		}
	}
	return true
}
