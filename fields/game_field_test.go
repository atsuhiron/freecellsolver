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
		Homes  map[uint8]*cells.HomeCell
		Frees  [consts.LenFre]*cells.FreeCell
		Fields [consts.LenFie]*cells.FieldCell
	}
	tests := []struct {
		name   string
		fields fields
		want   [consts.LenHash]uint64
	}{
		{
			name: "complex",
			fields: fields{
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

func TestGameField_IsFinished(t *testing.T) {
	type fields struct {
		Homes  map[uint8]*cells.HomeCell
		Frees  [consts.LenFre]*cells.FreeCell
		Fields [consts.LenFie]*cells.FieldCell
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "empty",
			fields: fields{
				Homes: map[uint8]*cells.HomeCell{
					suits[0]: {[]cards.Card{}, uint8(0)},
					suits[1]: {[]cards.Card{}, uint8(1)},
					suits[2]: {[]cards.Card{}, uint8(2)},
					suits[3]: {[]cards.Card{}, uint8(3)},
				},
				Frees: [consts.LenFre]*cells.FreeCell{
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
				},
				Fields: [consts.LenFie]*cells.FieldCell{
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
				},
			},
			want: false,
		},
		{
			name: "filled one home",
			fields: fields{
				Homes: map[uint8]*cells.HomeCell{
					suits[0]: {[]cards.Card{{Code: uint8(1)}, {Code: uint8(2)}, {Code: uint8(3)}, {Code: uint8(4)}, {Code: uint8(5)}, {Code: uint8(6)}, {Code: uint8(7)}, {Code: uint8(8)}, {Code: uint8(9)}, {Code: uint8(10)}, {Code: uint8(11)}, {Code: uint8(12)}, {Code: uint8(13)}}, uint8(0)},
					suits[1]: {[]cards.Card{}, uint8(1)},
					suits[2]: {[]cards.Card{}, uint8(2)},
					suits[3]: {[]cards.Card{}, uint8(3)},
				},
				Frees: [consts.LenFre]*cells.FreeCell{
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
				},
				Fields: [consts.LenFie]*cells.FieldCell{
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
				},
			},
			want: false,
		},
		{
			name: "filled all home",
			fields: fields{
				Homes: map[uint8]*cells.HomeCell{
					suits[0]: {[]cards.Card{{Code: uint8(1)}, {Code: uint8(2)}, {Code: uint8(3)}, {Code: uint8(4)}, {Code: uint8(5)}, {Code: uint8(6)}, {Code: uint8(7)}, {Code: uint8(8)}, {Code: uint8(9)}, {Code: uint8(10)}, {Code: uint8(11)}, {Code: uint8(12)}, {Code: uint8(13)}}, uint8(0)},
					suits[1]: {[]cards.Card{{Code: uint8(17)}, {Code: uint8(18)}, {Code: uint8(19)}, {Code: uint8(20)}, {Code: uint8(21)}, {Code: uint8(22)}, {Code: uint8(23)}, {Code: uint8(24)}, {Code: uint8(25)}, {Code: uint8(26)}, {Code: uint8(27)}, {Code: uint8(28)}, {Code: uint8(29)}}, uint8(1)},
					suits[2]: {[]cards.Card{{Code: uint8(33)}, {Code: uint8(34)}, {Code: uint8(35)}, {Code: uint8(36)}, {Code: uint8(37)}, {Code: uint8(38)}, {Code: uint8(39)}, {Code: uint8(40)}, {Code: uint8(41)}, {Code: uint8(42)}, {Code: uint8(43)}, {Code: uint8(44)}, {Code: uint8(45)}}, uint8(2)},
					suits[3]: {[]cards.Card{{Code: uint8(49)}, {Code: uint8(50)}, {Code: uint8(51)}, {Code: uint8(52)}, {Code: uint8(53)}, {Code: uint8(54)}, {Code: uint8(55)}, {Code: uint8(56)}, {Code: uint8(57)}, {Code: uint8(58)}, {Code: uint8(59)}, {Code: uint8(60)}, {Code: uint8(61)}}, uint8(3)},
				},
				Frees: [consts.LenFre]*cells.FreeCell{
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
				},
				Fields: [consts.LenFie]*cells.FieldCell{
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gf := GameField{
				Homes:  tt.fields.Homes,
				Frees:  tt.fields.Frees,
				Fields: tt.fields.Fields,
			}
			if got := gf.IsFinished(); got != tt.want {
				t.Errorf("IsFinished() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcFieldHash(t *testing.T) {
	type args struct {
		fields [consts.LenFie]*cells.FieldCell
	}
	tests := []struct {
		name string
		args args
		want [consts.MaxFieNum]uint64
	}{
		{
			name: "empty (nil)",
			args: args{
				fields: [consts.LenFie]*cells.FieldCell{},
			},
			want: [consts.MaxFieNum]uint64{
				0, 0, 0, 0, 0,
				0, 0, 0, 0, 0,
				0, 0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "empty (defined)",
			args: args{
				fields: [consts.LenFie]*cells.FieldCell{
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
				},
			},
			want: [consts.MaxFieNum]uint64{
				0, 0, 0, 0, 0,
				0, 0, 0, 0, 0,
				0, 0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "one card",
			args: args{
				fields: [consts.LenFie]*cells.FieldCell{
					{CardStack: []cards.Card{{Code: uint8(7)}}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
				},
			},
			want: [consts.MaxFieNum]uint64{
				7, 0, 0, 0, 0,
				0, 0, 0, 0, 0,
				0, 0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "one card (another order)",
			args: args{
				fields: [consts.LenFie]*cells.FieldCell{
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{{Code: uint8(7)}}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
				},
			},
			want: [consts.MaxFieNum]uint64{
				7, 0, 0, 0, 0,
				0, 0, 0, 0, 0,
				0, 0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "multi card, one field",
			args: args{
				fields: [consts.LenFie]*cells.FieldCell{
					{CardStack: []cards.Card{{Code: uint8(7)}, {Code: uint8(22)}}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
				},
			},
			want: [consts.MaxFieNum]uint64{
				7, 22, 0, 0, 0,
				0, 0, 0, 0, 0,
				0, 0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "multi card, one field (another order)",
			args: args{
				fields: [consts.LenFie]*cells.FieldCell{
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{{Code: uint8(7)}, {Code: uint8(22)}}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
				},
			},
			want: [consts.MaxFieNum]uint64{
				7, 22, 0, 0, 0,
				0, 0, 0, 0, 0,
				0, 0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "multi card, multi field",
			args: args{
				fields: [consts.LenFie]*cells.FieldCell{
					{CardStack: []cards.Card{{Code: uint8(7)}, {Code: uint8(22)}}},
					{CardStack: []cards.Card{{Code: uint8(35)}, {Code: uint8(18)}}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
				},
			},
			want: [consts.MaxFieNum]uint64{
				1827, 5650, 0, 0, 0,
				0, 0, 0, 0, 0,
				0, 0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "multi card, multi field (another order)",
			args: args{
				fields: [consts.LenFie]*cells.FieldCell{
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
			want: [consts.MaxFieNum]uint64{
				1827, 5650, 0, 0, 0,
				0, 0, 0, 0, 0,
				0, 0, 0, 0, 0,
				0, 0, 0, 0,
			},
		},
		{
			name: "multi card, one field (long sequence)",
			args: args{
				fields: [consts.LenFie]*cells.FieldCell{
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{{Code: uint8(13)}, {Code: uint8(28)}, {Code: uint8(11)}, {Code: uint8(26)}, {Code: uint8(9)}, {Code: uint8(24)}, {Code: uint8(7)}, {Code: uint8(22)}, {Code: uint8(5)}, {Code: uint8(20)}, {Code: uint8(3)}, {Code: uint8(18)}, {Code: uint8(1)}}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
				},
			},
			want: [consts.MaxFieNum]uint64{
				13, 28, 11, 26, 9,
				24, 7, 22, 5, 20,
				3, 18, 1, 0, 0,
				0, 0, 0, 0,
			},
		},
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
		frees [consts.LenFre]*cells.FreeCell
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "empty (nil)",
			args: args{
				frees: [consts.LenFre]*cells.FreeCell{},
			},
			want: uint64(0),
		},
		{
			name: "empty (defined)",
			args: args{
				frees: [consts.LenFre]*cells.FreeCell{
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
				},
			},
			want: uint64(0),
		},
		{
			name: "one card 1",
			args: args{
				frees: [consts.LenFre]*cells.FreeCell{
					{CardStack: []cards.Card{{Code: uint8(5)}}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
				},
			},
			want: uint64(5),
		},
		{
			name: "one card 1 (another order)",
			args: args{
				frees: [consts.LenFre]*cells.FreeCell{
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{{Code: uint8(5)}}},
					{CardStack: []cards.Card{}},
				},
			},
			want: uint64(5),
		},
		{
			name: "multi cards",
			args: args{
				frees: [consts.LenFre]*cells.FreeCell{
					{CardStack: []cards.Card{{Code: uint8(34)}}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{{Code: uint8(18)}}},
					{CardStack: []cards.Card{}},
				},
			},
			want: uint64(4642), // 0x00001222
		},
		{
			name: "multi cards (another order)",
			args: args{
				frees: [consts.LenFre]*cells.FreeCell{
					{CardStack: []cards.Card{{Code: uint8(18)}}},
					{CardStack: []cards.Card{{Code: uint8(34)}}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
				},
			},
			want: uint64(4642), // 0x00001222
		},
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
		homes map[uint8]*cells.HomeCell
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "empty (nil)",
			args: args{
				homes: map[uint8]*cells.HomeCell{
					suits[0]: {},
					suits[1]: {},
					suits[2]: {},
					suits[3]: {},
				},
			},
			want: uint64(0),
		},
		{
			name: "empty (defined)",
			args: args{
				homes: map[uint8]*cells.HomeCell{
					suits[0]: {[]cards.Card{}, uint8(0)},
					suits[1]: {[]cards.Card{}, uint8(1)},
					suits[2]: {[]cards.Card{}, uint8(2)},
					suits[3]: {[]cards.Card{}, uint8(3)},
				},
			},
			want: uint64(0),
		},
		{
			name: "one card 1",
			args: args{
				homes: map[uint8]*cells.HomeCell{
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
				homes: map[uint8]*cells.HomeCell{
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
				homes: map[uint8]*cells.HomeCell{
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
				homes: map[uint8]*cells.HomeCell{
					suits[0]: {},
					suits[1]: {},
					suits[2]: {},
					suits[3]: {CardStack: []cards.Card{{Code: uint8(49)}}},
				},
			},
			want: uint64(822083584), // 0x31000000
		},
		{
			name: "multi cards",
			args: args{
				homes: map[uint8]*cells.HomeCell{
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
		{
			name: "empty",
			iv:   indexValue64{},
			want: 0,
		},
		{
			name: "filled",
			iv: indexValue64{
				{uint64(1), uint64(35)},
				{uint64(2), uint64(38)},
			},
			want: 2,
		},
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
		{
			name: "1",
			iv: indexValue64{
				{uint64(1), uint64(35)},
				{uint64(2), uint64(38)},
			},
			args: args{0, 1},
			want: false,
		},
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
		want indexValue64
	}{
		{
			name: "1",
			iv: indexValue64{
				{uint64(1), uint64(35)},
				{uint64(2), uint64(38)},
			},
			args: args{0, 1},
			want: indexValue64{
				{uint64(2), uint64(38)},
				{uint64(1), uint64(35)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.iv.Swap(tt.args.i, tt.args.j)
			if !reflect.DeepEqual(tt.iv, tt.want) {
				t.Errorf("Swap() = %v, want %v", tt.iv, tt.want)
			}
		})
	}
}

func Test_calcMaxMovableCardNum(t *testing.T) {
	type args struct {
		free  int
		field int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0, 0",
			args: args{
				free:  0,
				field: 0,
			},
			want: 1,
		},
		{
			name: "0, 1",
			args: args{
				free:  0,
				field: 1,
			},
			want: 2,
		},
		{
			name: "1, 0",
			args: args{
				free:  1,
				field: 0,
			},
			want: 2,
		},
		{
			name: "2, 3",
			args: args{
				free:  2,
				field: 3,
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcMaxMovableCardNum(tt.args.free, tt.args.field); got != tt.want {
				t.Errorf("calcMaxMovableCardNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createGameField() GameField {
	return GameField{
		Homes: map[uint8]*cells.HomeCell{
			suits[0]: {
				SuitCode:  suits[0],
				CardStack: []cards.Card{},
			},
			suits[1]: {
				SuitCode:  suits[1],
				CardStack: []cards.Card{{Code: uint8(17)}},
			},
			suits[2]: {
				SuitCode:  suits[2],
				CardStack: []cards.Card{},
			},
			suits[3]: {
				SuitCode:  suits[3],
				CardStack: []cards.Card{},
			},
		},
		Frees: [consts.LenFre]*cells.FreeCell{
			{CardStack: []cards.Card{}},
			{CardStack: []cards.Card{}},
			{CardStack: []cards.Card{{Code: uint8(34)}}},
			{CardStack: []cards.Card{{Code: uint8(1)}}},
		},
		Fields: [consts.LenFie]*cells.FieldCell{
			{
				CardStack: []cards.Card{{Code: uint8(3)}, {Code: uint8(18)}},
			},
			{
				CardStack: []cards.Card{{Code: uint8(39)}, {Code: uint8(22)}},
			},
			{
				CardStack: []cards.Card{{Code: uint8(24)}},
			},
			{
				CardStack: []cards.Card{{Code: uint8(35)}},
			},
			{
				CardStack: []cards.Card{},
			},
			{
				CardStack: []cards.Card{},
			},
			{
				CardStack: []cards.Card{},
			},
			{
				CardStack: []cards.Card{},
			},
		},
	}
}

func TestGameField_move(t *testing.T) {
	type args struct {
		fieldTypeFrom string
		indexFrom     int
		fieldTypeTo   string
		indexTo       int
	}
	tests := []struct {
		name       string
		fields     GameField
		args       args
		wantFields GameField
		wantErr    bool
		wantErrMsg string
	}{
		{
			name:   "free to home",
			fields: createGameField(),
			args: args{
				fieldTypeFrom: "free",
				indexFrom:     3,
				fieldTypeTo:   "home",
				indexTo:       0,
			},
			wantFields: GameField{
				Homes: map[uint8]*cells.HomeCell{
					suits[0]: {
						SuitCode:  suits[0],
						CardStack: []cards.Card{{Code: uint8(1)}},
					},
					suits[1]: {
						SuitCode:  suits[1],
						CardStack: []cards.Card{{Code: uint8(17)}},
					},
					suits[2]: {
						SuitCode:  suits[2],
						CardStack: []cards.Card{},
					},
					suits[3]: {
						SuitCode:  suits[3],
						CardStack: []cards.Card{},
					},
				},
				Frees: [consts.LenFre]*cells.FreeCell{
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{{Code: uint8(34)}}},
					{CardStack: []cards.Card{}},
				},
				Fields: [consts.LenFie]*cells.FieldCell{
					{
						CardStack: []cards.Card{{Code: uint8(3)}, {Code: uint8(18)}},
					},
					{
						CardStack: []cards.Card{{Code: uint8(39)}, {Code: uint8(22)}},
					},
					{
						CardStack: []cards.Card{{Code: uint8(24)}},
					},
					{
						CardStack: []cards.Card{{Code: uint8(35)}},
					},
					{
						CardStack: []cards.Card{},
					},
					{
						CardStack: []cards.Card{},
					},
					{
						CardStack: []cards.Card{},
					},
					{
						CardStack: []cards.Card{},
					},
				},
			},
			wantErr:    false,
			wantErrMsg: "",
		},
		{
			name:   "field to home",
			fields: createGameField(),
			args: args{
				fieldTypeFrom: "field",
				indexFrom:     0,
				fieldTypeTo:   "home",
				indexTo:       1,
			},
			wantFields: GameField{
				Homes: map[uint8]*cells.HomeCell{
					suits[0]: {
						SuitCode:  suits[0],
						CardStack: []cards.Card{},
					},
					suits[1]: {
						SuitCode:  suits[1],
						CardStack: []cards.Card{{Code: uint8(17)}, {Code: uint8(18)}},
					},
					suits[2]: {
						SuitCode:  suits[2],
						CardStack: []cards.Card{},
					},
					suits[3]: {
						SuitCode:  suits[3],
						CardStack: []cards.Card{},
					},
				},
				Frees: [consts.LenFre]*cells.FreeCell{
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{}},
					{CardStack: []cards.Card{{Code: uint8(34)}}},
					{CardStack: []cards.Card{{Code: uint8(1)}}},
				},
				Fields: [consts.LenFie]*cells.FieldCell{
					{
						CardStack: []cards.Card{{Code: uint8(3)}},
					},
					{
						CardStack: []cards.Card{{Code: uint8(39)}, {Code: uint8(22)}},
					},
					{
						CardStack: []cards.Card{{Code: uint8(24)}},
					},
					{
						CardStack: []cards.Card{{Code: uint8(35)}},
					},
					{
						CardStack: []cards.Card{},
					},
					{
						CardStack: []cards.Card{},
					},
					{
						CardStack: []cards.Card{},
					},
					{
						CardStack: []cards.Card{},
					},
				},
			},
			wantErr:    false,
			wantErrMsg: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gf := &GameField{
				Homes:  tt.fields.Homes,
				Frees:  tt.fields.Frees,
				Fields: tt.fields.Fields,
			}

			err := gf.move(tt.args.fieldTypeFrom, tt.args.indexFrom, tt.args.fieldTypeTo, tt.args.indexTo)

			if (err != nil) != tt.wantErr {
				t.Errorf("move() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				if err.Error() != tt.wantErrMsg {
					t.Errorf("move() errorMsg = %v, wantErrMsg %v", err.Error(), tt.wantErrMsg)
				}
			} else {
				// check HomeCell equality
				for _, c := range suits {
					act := gf.Homes[c].CardStack
					exp := tt.wantFields.Homes[c].CardStack
					if !cells.EqualStack(&act, &exp) {
						t.Errorf("move() it does not match CardStack of Home[%v]", c)
					}
				}

				// check FreeCell equality
				for i := 0; i < consts.LenFre; i++ {
					act := gf.Frees[i].CardStack
					exp := tt.wantFields.Frees[i].CardStack
					if !cells.EqualStack(&act, &exp) {
						t.Errorf("move() it does not match CardStack of Free[%v]", i)
					}
				}

				// check FieldCell equality
				for i := 0; i < consts.LenFie; i++ {
					act := gf.Fields[i].CardStack
					exp := tt.wantFields.Fields[i].CardStack
					if !cells.EqualStack(&act, &exp) {
						t.Errorf("move() it does not match CardStack of Field[%v]", i)
					}
				}
			}
		})
	}
}
