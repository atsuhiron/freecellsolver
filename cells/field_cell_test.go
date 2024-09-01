package cells

import (
	"github.com/freecellsolver/cards"
	"reflect"
	"testing"
)

func TestFieldCell_CanPlace(t *testing.T) {
	type fields struct {
		CardStack []cards.Card
	}
	type args struct {
		card cards.Card
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Empty stack: OK",
			fields: fields{
				CardStack: []cards.Card{},
			},
			args: args{
				card: cards.Card{Code: uint8(1)},
			},
			want: true,
		},
		{
			name: "Filled stack: same color NG",
			fields: fields{
				CardStack: []cards.Card{
					{uint8(2)},
				},
			},
			args: args{
				card: cards.Card{Code: uint8(1)},
			},
			want: false,
		},
		{
			name: "Filled stack: different color OK",
			fields: fields{
				CardStack: []cards.Card{
					{uint8(18)},
				},
			},
			args: args{
				card: cards.Card{Code: uint8(1)},
			},
			want: true,
		},
		{
			name: "Filled stack: different color incorrect order NG",
			fields: fields{
				CardStack: []cards.Card{
					{uint8(19)},
				},
			},
			args: args{
				card: cards.Card{Code: uint8(1)},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fCell := FieldCell{
				CardStack: tt.fields.CardStack,
			}
			if got := fCell.CanPlace(tt.args.card); got != tt.want {
				t.Errorf("CanPlace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldCell_GetEndSeq(t *testing.T) {
	type fields struct {
		CardStack []cards.Card
	}
	tests := []struct {
		name   string
		fields fields
		want   []cards.Card
	}{
		{
			name: "Empty stack: return empty",
			fields: fields{
				CardStack: []cards.Card{},
			},
			want: []cards.Card{},
		},
		{
			name: "Single stack: return single card",
			fields: fields{
				CardStack: []cards.Card{
					{uint8(1)},
				},
			},
			want: []cards.Card{
				{uint8(1)},
			},
		},
		{
			name: "Continued and same color stack: return last",
			fields: fields{
				CardStack: []cards.Card{
					{uint8(6)},
					{uint8(5)},
					{uint8(4)},
					{uint8(3)},
				},
			},
			want: []cards.Card{
				{uint8(3)},
			},
		},
		{
			name: "Discontinued and same color stack: return last",
			fields: fields{
				CardStack: []cards.Card{
					{uint8(6)},
					{uint8(9)},
					{uint8(1)},
					{uint8(3)},
				},
			},
			want: []cards.Card{
				{uint8(3)},
			},
		},
		{
			name: "Continued and alternation color stack: return all",
			fields: fields{
				CardStack: []cards.Card{
					{uint8(6)},
					{uint8(21)},
					{uint8(4)},
					{uint8(19)},
				},
			},
			want: []cards.Card{
				{uint8(6)},
				{uint8(21)},
				{uint8(4)},
				{uint8(19)},
			},
		},
		{
			name: "Discontinued and alternation color stack: return last",
			fields: fields{
				CardStack: []cards.Card{
					{uint8(22)},
					{uint8(41)},
					{uint8(17)},
					{uint8(35)},
				},
			},
			want: []cards.Card{
				{uint8(35)},
			},
		},
		{
			name: "Partial continued and alternation color stack: return the part",
			fields: fields{
				CardStack: []cards.Card{
					{uint8(45)}, // ♣K
					{uint8(10)}, // ♠10
					{uint8(18)}, // ♥2
					{uint8(39)}, // ♣7
					{uint8(54)}, // ♦6
					{uint8(5)},  // ♠5
				},
			},
			want: []cards.Card{
				{uint8(39)}, // ♣7
				{uint8(54)}, // ♦6
				{uint8(5)},  // ♠5
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fCell := FieldCell{
				CardStack: tt.fields.CardStack,
			}
			if got := fCell.GetEndSeq(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEndSeq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldCell_Clone(t *testing.T) {
	type fields struct {
		CardStack []cards.Card
	}
	tests := []struct {
		name   string
		fields fields
		want   FieldCell
	}{
		{
			name: "empty stack",
			fields: fields{
				CardStack: []cards.Card{},
			},
			want: FieldCell{
				CardStack: []cards.Card{},
			},
		},
		{
			name: "single stack",
			fields: fields{
				CardStack: []cards.Card{{uint8(1)}},
			},
			want: FieldCell{
				CardStack: []cards.Card{{uint8(1)}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fCell := FieldCell{
				CardStack: tt.fields.CardStack,
			}
			got := fCell.Clone()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Clone() = %v, want %v", got, tt.want)
			}
			if &got.CardStack == &tt.fields.CardStack {
				t.Errorf("Cloned cell has same pointer")
			}
			if len(got.CardStack) > 0 {
				if &(got.CardStack[0]) == &(tt.fields.CardStack[0]) {
					t.Errorf("Cloned card has same pointer")
				}
			}
		})
	}
}

func TestFieldCell_Place(t *testing.T) {
	type fields struct {
		CardStack []cards.Card
	}
	type args struct {
		seq *[]cards.Card
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []cards.Card
	}{
		{
			name: "empty: place empty stack",
			fields: fields{
				CardStack: []cards.Card{},
			},
			args: args{
				seq: &[]cards.Card{},
			},
			want: []cards.Card{},
		},
		{
			name: "empty: place 1 card",
			fields: fields{
				CardStack: []cards.Card{},
			},
			args: args{
				seq: &[]cards.Card{{Code: uint8(1)}},
			},
			want: []cards.Card{{Code: uint8(1)}},
		},
		{
			name: "filled: place 1 card",
			fields: fields{
				CardStack: []cards.Card{{uint8(5)}},
			},
			args: args{
				seq: &[]cards.Card{{Code: uint8(20)}},
			},
			want: []cards.Card{
				{Code: uint8(5)},
				{Code: uint8(20)},
			},
		},
		{
			name: "filled: place multi cards",
			fields: fields{
				CardStack: []cards.Card{{uint8(5)}},
			},
			args: args{
				seq: &[]cards.Card{
					{Code: uint8(20)},
					{Code: uint8(3)},
				},
			},
			want: []cards.Card{
				{Code: uint8(5)},
				{Code: uint8(20)},
				{Code: uint8(3)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fCell := &FieldCell{
				CardStack: tt.fields.CardStack,
			}

			fCell.Place(tt.args.seq)
			if !EqualStack(&(fCell.CardStack), &(tt.want)) {
				t.Errorf("Place() = %v, want %v", fCell.CardStack, tt.want)
			}
		})
	}
}
