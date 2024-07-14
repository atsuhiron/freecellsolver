package fields

import (
	"github.com/freecellsolver/cells"
	"github.com/freecellsolver/consts"
)

type GameField struct {
	Homes  map[uint8]cells.HomeCell
	Frees  [consts.LenFre]cells.FreeCell
	Fields [consts.LenFie]cells.FieldCell
}

func (gf GameField) GetHashCode() int {
	homeCode := uint32(0)
	for _, sc := range suits {
		stack := gf.Homes[sc].CardStack
		homeCode += uint32(stack[len(stack)-1].Code << (8 * sc))
	}

	return 0
}
