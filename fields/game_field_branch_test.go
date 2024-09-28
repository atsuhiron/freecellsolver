package fields

import (
	"github.com/freecellsolver/cards"
	"github.com/freecellsolver/cells"
	"github.com/freecellsolver/consts"
	"reflect"
	"testing"
)

func TestBranchLattice_Len(t *testing.T) {
	tests := []struct {
		name string
		bl   BranchLattice
		want int
	}{
		{
			name: "empty",
			bl:   BranchLattice{},
			want: 0,
		},
		{
			name: "filled",
			bl: BranchLattice{
				GameFieldBranch{
					GF:   GameField{},
					Cost: 0,
				},
				GameFieldBranch{
					GF:   GameField{},
					Cost: 1,
				},
				GameFieldBranch{
					GF:   GameField{},
					Cost: 2,
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bl.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBranchLattice_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		bl   BranchLattice
		args args
		want bool
	}{
		{
			name: "true",
			bl: BranchLattice{
				GameFieldBranch{
					GF:   GameField{},
					Cost: 0,
				},
				GameFieldBranch{
					GF:   GameField{},
					Cost: 1,
				},
			},
			args: args{0, 1},
			want: true,
		},
		{
			name: "false",
			bl: BranchLattice{
				GameFieldBranch{
					GF:   GameField{},
					Cost: 1,
				},
				GameFieldBranch{
					GF:   GameField{},
					Cost: -1,
				},
			},
			args: args{0, 1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bl.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBranchLattice_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		bl   BranchLattice
		args args
		want BranchLattice
	}{
		{
			name: "1",
			bl: BranchLattice{
				GameFieldBranch{
					GF:   GameField{},
					Cost: 0,
				},
				GameFieldBranch{
					GF:   GameField{},
					Cost: 1,
				},
			},
			args: args{0, 1},
			want: BranchLattice{
				GameFieldBranch{
					GF:   GameField{},
					Cost: 1,
				},
				GameFieldBranch{
					GF:   GameField{},
					Cost: 0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bl.Swap(tt.args.i, tt.args.j)
			if !reflect.DeepEqual(tt.bl, tt.want) {
				t.Errorf("Swap() = %v, want %v", tt.bl, tt.want)
			}
		})
	}
}

func TestGameFieldBranch_CalcHashCode(t *testing.T) {
	type fields struct {
		GF   GameField
		Cost int8
	}
	tests := []struct {
		name   string
		fields fields
		want   [consts.LenHash]uint64
	}{
		{
			name: "complex",
			fields: fields{
				GF: GameField{
					Homes: map[uint8]*cells.HomeCell{
						suits[0]: {CardStack: []cards.Card{{Code: uint8(1)}}},
						suits[1]: {},
						suits[2]: {CardStack: []cards.Card{{Code: uint8(33)}, {Code: uint8(34)}}},
						suits[3]: {},
					},
					Frees: [consts.LenFre]*cells.FreeCell{
						{CardStack: []cards.Card{{Code: uint8(34)}}},
						{CardStack: []cards.Card{}},
						{CardStack: []cards.Card{{Code: uint8(18)}}},
						{CardStack: []cards.Card{}},
					},
					Fields: [consts.LenFie]*cells.FieldCell{
						{CardStack: []cards.Card{}},
						{CardStack: []cards.Card{{Code: uint8(35)}, {Code: uint8(18)}}},
						{CardStack: []cards.Card{}},
						{CardStack: []cards.Card{{Code: uint8(7)}, {Code: uint8(22)}}},
						{CardStack: []cards.Card{}},
						{CardStack: []cards.Card{}},
						{CardStack: []cards.Card{}},
						{CardStack: []cards.Card{}},
					},
				},
				Cost: 1,
			},
			want: [consts.LenHash]uint64{
				9570153503134242, // 0x0022000100001222  (2228225 << 32 + 4642)
				1827, 5650, 0, 0, 0,
				0, 0, 0, 0, 0,
				0, 0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gfb := &GameFieldBranch{
				GF:   tt.fields.GF,
				Cost: tt.fields.Cost,
			}
			if got := gfb.CalcHashCode(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalcHashCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
