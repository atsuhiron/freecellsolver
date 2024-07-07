package cells

import (
	"github.com/freecellsolver/cards"
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
				cardStack: tt.fields.cardStack,
			}
			if got := fCell.CanPlace(tt.args.card); got != tt.want {
				t.Errorf("CanPlace() = %v, want %v", got, tt.want)
			}
		})
	}
}
