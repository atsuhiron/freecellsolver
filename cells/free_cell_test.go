package cells

import (
	"github.com/freecellsolver/cards"
	"reflect"
	"testing"
)

func TestFreeCell_CanPlace(t *testing.T) {
	type fields struct {
		cardStack []cards.Card
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
			name:   "filled",
			fields: fields{[]cards.Card{{uint8(1)}}},
			args:   args{card: cards.Card{Code: uint8(2)}},
			want:   false,
		},
		{
			name:   "empty",
			fields: fields{[]cards.Card{}},
			args:   args{card: cards.Card{Code: uint8(2)}},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fCell := FreeCell{
				CardStack: tt.fields.cardStack,
			}
			if got := fCell.CanPlace(tt.args.card); got != tt.want {
				t.Errorf("CanPlace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFreeCell_GetEndSeq(t *testing.T) {
	type fields struct {
		CardStack []cards.Card
	}
	type args struct {
		onlyLast bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []cards.Card
	}{
		{
			name:   "empty: return empty",
			fields: fields{[]cards.Card{}},
			args:   args{onlyLast: true},
			want:   []cards.Card{},
		},
		{
			name:   "filled: return it",
			fields: fields{[]cards.Card{{uint8(1)}}},
			args:   args{onlyLast: true},
			want:   []cards.Card{{uint8(1)}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fCell := FreeCell{
				CardStack: tt.fields.CardStack,
			}
			if got := fCell.GetEndSeq(tt.args.onlyLast); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEndSeq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFreeCell_Clone(t *testing.T) {
	type fields struct {
		CardStack []cards.Card
	}
	tests := []struct {
		name   string
		fields fields
		want   *FreeCell
	}{
		{
			name: "empty stack",
			fields: fields{
				CardStack: []cards.Card{},
			},
			want: &FreeCell{
				CardStack: []cards.Card{},
			},
		},
		{
			name: "single stack",
			fields: fields{
				CardStack: []cards.Card{{uint8(1)}},
			},
			want: &FreeCell{
				CardStack: []cards.Card{{uint8(1)}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fCell := FreeCell{
				CardStack: tt.fields.CardStack,
			}
			got := fCell.Clone()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Clone() = %v, want %v", *got, *tt.want)
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

func TestFreeCell_Place(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fCell := &FreeCell{
				CardStack: tt.fields.CardStack,
			}

			fCell.Place(tt.args.seq)
			if !EqualStack(&(fCell.CardStack), &(tt.want)) {
				t.Errorf("Place() = %v, want %v", fCell.CardStack, tt.want)
			}
		})
	}
}
