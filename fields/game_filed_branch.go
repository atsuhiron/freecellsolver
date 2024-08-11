package fields

import "github.com/freecellsolver/consts"

type GameFieldBranch struct {
	GF   GameField
	Cost int
}

func (gfb GameFieldBranch) CalcHashCode() [consts.LenHash]uint64 {
	return gfb.GF.CalcHashCode()
}
