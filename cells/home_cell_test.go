package cells

import (
	"github.com/freecellsolver/cards"
	"testing"
)

func TestHomeCell_CanPlace(t *testing.T) {
	type fields struct {
		CardStack []cards.Card
		SuitCode  uint8
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
			name: "CanPlace different color",
			fields: fields{
				CardStack: []cards.Card{
					{uint8(1)},
					{uint8(2)},
					{uint8(3)},
				},
				SuitCode: uint8(0),
			},
			args: args{card: cards.Card{Code: uint8(33)}},
			want: false,
		},
		{
			name: "CanPlace jump number",
			fields: fields{
				CardStack: []cards.Card{
					{uint8(1)},
					{uint8(2)},
				},
				SuitCode: uint8(0),
			},
			args: args{card: cards.Card{Code: uint8(5)}},
			want: false,
		},
		{
			name: "CanPlace correct order",
			fields: fields{
				CardStack: []cards.Card{
					{uint8(1)},
					{uint8(2)},
				},
				SuitCode: uint8(0),
			},
			args: args{card: cards.Card{Code: uint8(3)}},
			want: true,
		},
		{
			name: "CanPlace correct order first",
			fields: fields{
				CardStack: []cards.Card{},
				SuitCode:  uint8(0),
			},
			args: args{card: cards.Card{Code: uint8(1)}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hCell := HomeCell{
				CardStack: tt.fields.CardStack,
				SuitCode:  tt.fields.SuitCode,
			}
			if got := hCell.CanPlace(tt.args.card); got != tt.want {
				t.Errorf("CanPlace() = %v, want %v", got, tt.want)
			}
		})
	}
}
