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
	Homes  map[uint8]*cells.HomeCell
	Frees  [consts.LenFre]*cells.FreeCell
	Fields [consts.LenFie]*cells.FieldCell
}

func (gf *GameField) CalcHashCode() [consts.LenHash]uint64 {
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

func (gf *GameField) IsFinished() bool {
	for _, sc := range suits {
		if len(gf.Homes[sc].CardStack) < 13 {
			return false
		}
	}
	return true
}

func (gf *GameField) GetBranch() ([]GameFieldBranch, error) {
	emptyFieldNum := 0
	for i := 0; i < consts.LenFie; i++ {
		if len(gf.Fields[i].CardStack) == 0 {
			emptyFieldNum++
		}
	}
	emptyFreeNum := 0
	for i := 0; i < consts.LenFre; i++ {
		if len(gf.Frees[i].CardStack) == 0 {
			emptyFreeNum++
		}
	}
	maxMovableCardNum := calcMaxMovableCardNum(emptyFreeNum, emptyFieldNum)

	branches := make([]GameFieldBranch, 0, consts.LenFre+consts.LenFie)

	// from field
	for iSrc := 0; iSrc < consts.LenFie; iSrc++ {
		seq := gf.Fields[iSrc].GetEndSeq(false)
		seqSize := len(seq)
		if seqSize == 0 {
			continue
		}
		if seqSize > maxMovableCardNum {
			continue
		}

		// to home
		last := seq[seqSize-1]
		suitCode := last.GetSuitCode()
		if gf.Homes[suitCode].CanPlace(last) {
			cloned := gf.clone()
			err := cloned.move("field", iSrc, "home", int(suitCode))
			if err != nil {
				return branches, err
			}
			branches = append(branches, GameFieldBranch{cloned, -1})
		}

		// to field
		movedEmptyField := len(gf.Fields[iSrc].CardStack) == len(seq)
		for iTgt := 0; iTgt < consts.LenFie; iTgt++ {
			if iTgt == iSrc {
				// Move to self
				continue
			}

			if len(gf.Fields[iTgt].CardStack) == 0 && !movedEmptyField {
				// If target field is empty, it is allowed only first empty field
				movedEmptyField = true
			} else if len(gf.Fields[iTgt].CardStack) == 0 && movedEmptyField {
				// To avoid duplication
				continue
			}

			if gf.Fields[iTgt].CanPlace(seq[0]) {
				cloned := gf.clone()
				err := cloned.move("field", iSrc, "field", iTgt)
				if err != nil {
					return branches, err
				}
				branches = append(branches, GameFieldBranch{cloned, 0})
			}
		}

		// to free
		for iTgt := 0; iTgt < consts.LenFre; iTgt++ {
			if gf.Frees[iTgt].CanPlace(last) {
				cloned := gf.clone()
				err := cloned.move("field", iSrc, "free", iTgt)
				if err != nil {
					return branches, err
				}
				branches = append(branches, GameFieldBranch{cloned, 1})
				break // To avoid duplication
			}
		}
	}

	// from free
	for iSrc := 0; iSrc < consts.LenFre; iSrc++ {
		seq := gf.Frees[iSrc].GetEndSeq(true)
		if len(seq) == 0 {
			continue
		}

		// to home
		last := seq[0] // len(seq) == 1
		suitCode := last.GetSuitCode()
		if gf.Homes[suitCode].CanPlace(last) {
			cloned := gf.clone()
			err := cloned.move("free", iSrc, "home", int(suitCode))
			if err != nil {
				return branches, err
			}
			branches = append(branches, GameFieldBranch{cloned, -2})
		}

		// to field
		for iTgt := 0; iTgt < consts.LenFie; iTgt++ {
			if gf.Fields[iTgt].CanPlace(seq[0]) {
				cloned := gf.clone()
				err := cloned.move("free", iSrc, "field", iTgt)
				if err != nil {
					return branches, err
				}
				branches = append(branches, GameFieldBranch{cloned, 0})
			}
		}
	}
	return branches, nil
}

func calcMaxMovableCardNum(free int, field int) int {
	return (field + 1) * (free + 1)
}

func calcHomeHash(homes map[uint8]*cells.HomeCell) uint64 {
	homeCode := uint64(0)
	for _, sc := range suits {
		stack := homes[sc].CardStack
		if len(stack) != 0 {
			homeCode += uint64(stack[len(stack)-1].Code) << (8 * sc)
		}
	}
	return homeCode
}

func calcFreeHash(frees [consts.LenFre]*cells.FreeCell) uint64 {
	freeCardCodes := make([]uint64, consts.LenFre)
	for i, cell := range frees {
		if cell == nil {
			// nil
			freeCardCodes[i] = uint64(0)
		} else if len(cell.CardStack) == 0 {
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

func calcFieldHash(fields [consts.LenFie]*cells.FieldCell) [consts.MaxFieNum]uint64 {
	fieldCardCodes := make(indexValue64, consts.LenFie)
	for i, cell := range fields {
		fieldCardCodes[i][0] = uint64(i)
		if cell == nil {
			fieldCardCodes[i][1] = uint64(0)
			fields[i] = &(cells.FieldCell{})
		} else if len(cell.CardStack) == 0 {
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

func (gf *GameField) clone() GameField {
	homes := make(map[uint8]*cells.HomeCell)
	for suit, cell := range gf.Homes {
		homes[suit] = cell.Clone()
	}

	frees := [consts.LenFre]*cells.FreeCell{}
	for i := range gf.Frees {
		frees[i] = gf.Frees[i].Clone()
	}

	fields := [consts.LenFie]*cells.FieldCell{}
	for i := range gf.Fields {
		fields[i] = gf.Fields[i].Clone()
	}

	return GameField{
		Homes:  homes,
		Frees:  frees,
		Fields: fields,
	}
}

func (gf *GameField) move(fieldTypeFrom string, indexFrom int, fieldTypeTo string, indexTo int) error {
	if !strings.EqualFold(fieldTypeFrom, "free") && !strings.EqualFold(fieldTypeFrom, "field") {
		return fmt.Errorf("invalid fieldTypeFrom %v", fieldTypeFrom)
	}
	if !strings.EqualFold(fieldTypeTo, "home") && !strings.EqualFold(fieldTypeTo, "free") && !strings.EqualFold(fieldTypeTo, "field") {
		return fmt.Errorf("invalid fieldTypeTo %v", fieldTypeTo)
	}
	moveOnlyLast := strings.EqualFold(fieldTypeTo, "home") || strings.EqualFold(fieldTypeTo, "free")

	// cut
	var seq []cards.Card
	if strings.EqualFold(fieldTypeFrom, "free") {
		seq = make([]cards.Card, 0, 1) // 常に長さが1なので、cap=1
		c := gf.Frees[indexFrom]
		seq = c.GetEndSeq(moveOnlyLast)
		if err := c.RemoveEndSeq(len(seq)); err != nil {
			return err
		}
	} else {
		// Field
		seq = make([]cards.Card, 0, 12) // 長さは12以下であるので cap=12
		c := gf.Fields[indexFrom]
		seq = c.GetEndSeq(moveOnlyLast)
		if err := c.RemoveEndSeq(len(seq)); err != nil {
			return err
		}
	}

	// paste
	if strings.EqualFold(fieldTypeTo, "free") {
		gf.Frees[indexTo].Place(&seq)
	} else if strings.EqualFold(fieldTypeTo, "field") {
		gf.Fields[indexTo].Place(&seq)
	} else {
		// Home
		c := gf.Homes[uint8(indexTo)]
		c.Place(&seq)
	}

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
