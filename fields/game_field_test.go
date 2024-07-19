package fields

import (
	"github.com/freecellsolver/cards"
	"github.com/freecellsolver/cells"
	"github.com/freecellsolver/consts"
	"reflect"
	"testing"
)

func TestGameField_CalcHashCode(t *testing.T) {
	type fields struct {
		Homes  map[uint8]cells.HomeCell
		Frees  [consts.LenFre]cells.FreeCell
		Fields [consts.LenFie]cells.FieldCell
	}
	tests := []struct {
		name   string
		fields fields
		want   [consts.LenHash]uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gf := GameField{
				Homes:  tt.fields.Homes,
				Frees:  tt.fields.Frees,
				Fields: tt.fields.Fields,
			}
			if got := gf.CalcHashCode(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalcHashCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcFieldHash(t *testing.T) {
	type args struct {
		fields [consts.LenFie]cells.FieldCell
	}
	tests := []struct {
		name string
		args args
		want [consts.MaxFieNum]uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcFieldHash(tt.args.fields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcFieldHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcFreeHash(t *testing.T) {
	type args struct {
		frees [consts.LenFre]cells.FreeCell
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcFreeHash(tt.args.frees); got != tt.want {
				t.Errorf("calcFreeHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcHomeHash(t *testing.T) {
	type args struct {
		homes map[uint8]cells.HomeCell
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "empty",
			args: args{
				homes: map[uint8]cells.HomeCell{
					suits[0]: {},
					suits[1]: {},
					suits[2]: {},
					suits[3]: {},
				},
			},
			want: uint64(0),
		},
		{
			name: "one card 1",
			args: args{
				homes: map[uint8]cells.HomeCell{
					suits[0]: {CardStack: []cards.Card{{Code: uint8(1)}}},
					suits[1]: {},
					suits[2]: {},
					suits[3]: {},
				},
			},
			want: uint64(1), // 0x00000001
		},
		{
			name: "one card 2",
			args: args{
				homes: map[uint8]cells.HomeCell{
					suits[0]: {},
					suits[1]: {CardStack: []cards.Card{{Code: uint8(17)}}},
					suits[2]: {},
					suits[3]: {},
				},
			},
			want: uint64(4352), // 0x00001100
		},
		{
			name: "one card 3",
			args: args{
				homes: map[uint8]cells.HomeCell{
					suits[0]: {},
					suits[1]: {},
					suits[2]: {CardStack: []cards.Card{{Code: uint8(33)}}},
					suits[3]: {},
				},
			},
			want: uint64(2162688), // 0x00210000
		},
		{
			name: "one card 4",
			args: args{
				homes: map[uint8]cells.HomeCell{
					suits[0]: {},
					suits[1]: {},
					suits[2]: {},
					suits[3]: {CardStack: []cards.Card{{Code: uint8(49)}}},
				},
			},
			want: uint64(822083584), // 0x31000000
		},
		{
			name: "multi card",
			args: args{
				homes: map[uint8]cells.HomeCell{
					suits[0]: {CardStack: []cards.Card{{Code: uint8(1)}}},
					suits[1]: {},
					suits[2]: {CardStack: []cards.Card{{Code: uint8(33)}, {Code: uint8(34)}}},
					suits[3]: {},
				},
			},
			want: uint64(2228225), // 0x00220001
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcHomeHash(tt.args.homes); got != tt.want {
				t.Errorf("calcHomeHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_indexValue64_Len(t *testing.T) {
	tests := []struct {
		name string
		iv   indexValue64
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.iv.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_indexValue64_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		iv   indexValue64
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.iv.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_indexValue64_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		iv   indexValue64
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.iv.Swap(tt.args.i, tt.args.j)
		})
	}
}
