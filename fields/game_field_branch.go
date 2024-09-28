package fields

import "github.com/freecellsolver/consts"

type GameFieldBranch struct {
	GF   GameField
	Cost int8
}

func (gfb *GameFieldBranch) CalcHashCode() [consts.LenHash]uint64 {
	return gfb.GF.CalcHashCode()
}

type BranchLattice []GameFieldBranch

func (bl *BranchLattice) Len() int           { return len(*bl) }
func (bl *BranchLattice) Less(i, j int) bool { return (*bl)[i].Cost < (*bl)[j].Cost }
func (bl *BranchLattice) Swap(i, j int)      { (*bl)[i], (*bl)[j] = (*bl)[j], (*bl)[i] }
