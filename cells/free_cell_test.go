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
	tests := []struct {
		name   string
		fields fields
		want   []cards.Card
	}{
		{
			name:   "empty: return empty",
			fields: fields{[]cards.Card{}},
			want:   []cards.Card{},
		},
		{
			name:   "filled: return it",
			fields: fields{[]cards.Card{{uint8(1)}}},
			want:   []cards.Card{{uint8(1)}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fCell := FreeCell{
				CardStack: tt.fields.CardStack,
			}
			if got := fCell.GetEndSeq(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEndSeq() = %v, want %v", got, tt.want)
			}
		})
	}
}
