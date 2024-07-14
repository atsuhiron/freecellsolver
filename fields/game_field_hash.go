package fields

import "github.com/freecellsolver/consts"

type GameFieldHash struct {
	HomeCode   uint32
	FreeCode   uint32
	FieldCodes [consts.MaxFieNum]uint64
}
