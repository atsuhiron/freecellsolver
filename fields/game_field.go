package fields

import (
	"fmt"
	"github.com/freecellsolver/cards"
	"github.com/freecellsolver/cells"
	"github.com/freecellsolver/consts"
	"slices"
	"sort"
	"strings"
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

func (gf GameField) IsFinished() bool {
	for _, sc := range suits {
		if len(gf.Homes[sc].CardStack) < 13 {
			return false
		}
	}
	return true
}

func (gf GameField) GetBranch() []GameFieldBranch {
	emptyFieldMum := 0
	for i := 0; i < consts.LenFie; i++ {
		if len(gf.Fields[i].CardStack) == 0 {
			emptyFieldMum++
		}
	}

	branch := make([]GameFieldBranch, 0, consts.LenFre+consts.LenFie)
	for iSrc := 0; iSrc < consts.LenFie; iSrc++ {
		seq := gf.Fields[iSrc].GetEndSeq()
		if len(seq) == 0 {
			continue
		}

		for iTgt := 0; iTgt < consts.LenFie; iTgt++ {
			if iTgt == iSrc {
				// Move to self
				continue
			}

			if gf.Fields[iTgt].CanPlace(seq[0]) {
				// TODO: make branch!
				branch = append(branch)
			}
		}
	}
	return branch // TODO: Implement
}

func calcMaxMovableCardNum(free int, field int) int {
	return (field + 1) * (free + 1)
}

func calcHomeHash(homes map[uint8]cells.HomeCell) uint64 {
	// TODO: ポインタにする
	homeCode := uint64(0)
	for _, sc := range suits {
		stack := homes[sc].CardStack
		if len(stack) != 0 {
			homeCode += uint64(stack[len(stack)-1].Code) << (8 * sc)
		}
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
	slices.Reverse(freeCardCodes)
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
			// As long as the patterns are the same field,
			// the order does not matter, so only the first card is checked.
			fieldCardCodes[i][1] = uint64(cell.CardStack[0].Code)
		}
	}
	sort.Sort(fieldCardCodes)

	sortedField := [consts.LenFie][]cards.Card{}
	for i, ivp := range fieldCardCodes {
		sortedField[i] = fields[ivp[0]].CardStack
	}

	fieldCodes := [consts.MaxFieNum]uint64{}
	for j := range fieldCodes {
		fieldCode := uint64(0)
		for i, column := range sortedField {
			if len(column) > j {
				fieldCode += uint64(column[j].Code) << (8 * i)
			}
		}
		fieldCodes[j] = fieldCode
	}

	return fieldCodes
}

func (gf GameField) clone() GameField {
	homes := make(map[uint8]cells.HomeCell)
	for suit, cell := range gf.Homes {
		homes[suit] = cell.Clone()
	}

	frees := [consts.LenFre]cells.FreeCell{}
	for i := range gf.Frees {
		frees[i] = gf.Frees[i].Clone()
	}

	fields := [consts.LenFie]cells.FieldCell{}
	for i := range gf.Fields {
		fields[i] = gf.Fields[i].Clone()
	}

	return GameField{
		Homes:  homes,
		Frees:  frees,
		Fields: fields,
	}
}

func (gf GameField) move(fieldTypeFrom string, indexFrom int, fieldTypeTo string, indexTo int) error {
	if !strings.EqualFold(fieldTypeFrom, "home") || !strings.EqualFold(fieldTypeFrom, "free") || !strings.EqualFold(fieldTypeFrom, "field") {
		return fmt.Errorf("invalid fieldTypeFrom %v", fieldTypeFrom)
	}
	if !strings.EqualFold(fieldTypeTo, "home") || !strings.EqualFold(fieldTypeTo, "free") || !strings.EqualFold(fieldTypeTo, "field") {
		return fmt.Errorf("invalid fieldTypeTo %v", fieldTypeTo)
	}
	// TODO: implement
	return nil
}

type indexValue64 [][2]uint64

func (iv indexValue64) Len() int      { return len(iv) }
func (iv indexValue64) Swap(i, j int) { iv[i], iv[j] = iv[j], iv[i] }
func (iv indexValue64) Less(i, j int) bool {
	// To implement it as its name(Less) suggests, the slice need to be reverse.
	// To avoid this step, the inequality signs are reversed (`<` -> `<`).
	return iv[i][1] > iv[j][1]
}
