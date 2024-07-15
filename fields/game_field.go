package fields

import (
	"github.com/freecellsolver/cards"
	"github.com/freecellsolver/cells"
	"github.com/freecellsolver/consts"
	"slices"
	"sort"
)

type GameField struct {
	Homes  map[uint8]cells.HomeCell
	Frees  [consts.LenFre]cells.FreeCell
	Fields [consts.LenFie]cells.FieldCell
}

func (gf GameField) CalcHashCode() [consts.LenHash]uint64 {
	homeCode := calcHomeHash(gf.Homes)
	freeCode := calcFreeHash(gf.Frees)
	fieldCodes := calcFieldHash(gf.Fields)

	hashCodes := [consts.LenHash]uint64{}
	hashCodes[0] = (homeCode << 32) + freeCode
	for i := 1; i < consts.LenHash; i++ {
		hashCodes[i] = fieldCodes[i-1]
	}
	return hashCodes
}

func calcHomeHash(homes map[uint8]cells.HomeCell) uint64 {
	// TODO: ポインタにする
	homeCode := uint64(0)
	for _, sc := range suits {
		stack := homes[sc].CardStack
		homeCode += uint64(stack[len(stack)-1].Code << (8 * sc))
	}
	return homeCode
}

func calcFreeHash(frees [consts.LenFre]cells.FreeCell) uint64 {
	freeCardCodes := make([]uint64, consts.LenFre)
	for i, cell := range frees {
		if len(cell.CardStack) == 0 {
			// empty
			freeCardCodes[i] = uint64(0)
		} else {
			// filled
			freeCardCodes[i] = uint64(cell.CardStack[0].Code)
		}
	}
	slices.Sort(freeCardCodes)
	freeCode := uint64(0)
	for i, cardCode := range freeCardCodes {
		freeCode += cardCode << (8 * i)
	}

	return freeCode
}

func calcFieldHash(fields [consts.LenFie]cells.FieldCell) [consts.MaxFieNum]uint64 {
	fieldCardCodes := make(indexValue64, consts.LenFie)
	for i, cell := range fields {
		fieldCardCodes[i][0] = uint64(i)

		if len(cell.CardStack) == 0 {
			fieldCardCodes[i][1] = uint64(0)
		} else {
			fieldCardCodes[i][1] = uint64(cell.CardStack[0].Code)
		}
	}
	sort.Sort(fieldCardCodes)

	sortedField := [consts.LenFie][]cards.Card{}
	for i, ivp := range fieldCardCodes {
		sortedField[i] = fields[ivp[1]].CardStack
	}

	fieldCodes := [consts.MaxFieNum]uint64{}
	for j, _ := range fieldCodes {
		fieldCode := uint64(0)
		for i, column := range sortedField {
			if len(column) > j {
				fieldCode += uint64(column[i].Code) << (8 * i)
			}
		}
		fieldCodes[j] = fieldCode
	}

	return fieldCodes
}

type indexValue64 [][2]uint64

func (iv indexValue64) Len() int           { return len(iv) }
func (iv indexValue64) Swap(i, j int)      { iv[i], iv[j] = iv[j], iv[i] }
func (iv indexValue64) Less(i, j int) bool { return iv[i][1] < iv[j][1] }
