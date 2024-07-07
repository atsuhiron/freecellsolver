package cells

import (
	"github.com/freecellsolver/cards"
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
					{uint8(34)},
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
					{uint8(35)},
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
