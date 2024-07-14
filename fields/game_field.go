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
