package cells

import (
	"github.com/freecellsolver/cards"
	"reflect"
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
			name: "different color",
			fields: fields{
				CardStack: []cards.Card{
					{uint8(1)},
					{uint8(2)},
					{uint8(3)},
				},
				SuitCode: uint8(0),
			},
			args: args{card: cards.Card{Code: uint8(17)}},
			want: false,
		},
		{
			name: "jump number",
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
			name: "correct order",
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
			name: "correct order first",
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

func TestHomeCell_GetEndSeq(t *testing.T) {
	type fields struct {
		CardStack []cards.Card
		SuitCode  uint8
	}
	tests := []struct {
		name   string
		fields fields
		want   []cards.Card
	}{
		{
			name: "always returns empty 1",
			fields: fields{
				CardStack: []cards.Card{},
				SuitCode:  uint8(0),
			},
			want: []cards.Card{},
		},
		{
			name: "always returns empty 2",
			fields: fields{
				CardStack: []cards.Card{
					{uint8(1)},
					{uint8(2)},
					{uint8(3)},
					{uint8(4)},
				},
				SuitCode: uint8(0),
			},
			want: []cards.Card{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hCell := HomeCell{
				CardStack: tt.fields.CardStack,
				SuitCode:  tt.fields.SuitCode,
			}
			if got := hCell.GetEndSeq(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEndSeq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHomeCell_Clone(t *testing.T) {
	type fields struct {
		CardStack []cards.Card
		SuitCode  uint8
	}
	tests := []struct {
		name   string
		fields fields
		want   HomeCell
	}{
		{
			name: "empty stack",
			fields: fields{
				CardStack: []cards.Card{},
				SuitCode:  uint8(0),
			},
			want: HomeCell{
				CardStack: []cards.Card{},
				SuitCode:  uint8(0),
			},
		},
		{
			name: "single stack",
			fields: fields{
				CardStack: []cards.Card{{uint8(1)}},
				SuitCode:  uint8(0),
			},
			want: HomeCell{
				CardStack: []cards.Card{{uint8(1)}},
				SuitCode:  uint8(0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hCell := HomeCell{
				CardStack: tt.fields.CardStack,
				SuitCode:  tt.fields.SuitCode,
			}
			got := hCell.Clone()
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

func TestHomeCell_RemoveEndSeq(t *testing.T) {
	type fields struct {
		CardStack []cards.Card
		SuitCode  uint8
	}
	type args struct {
		removeNum int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "empty: move 0 OK",
			fields: fields{
				CardStack: []cards.Card{},
			},
			args:    args{removeNum: 0},
			wantErr: false,
		},
		{
			name: "empty: move 1 error",
			fields: fields{
				CardStack: []cards.Card{},
			},
			args:    args{removeNum: 1},
			wantErr: true,
		},
		{
			name: "filled: move 0 OK",
			fields: fields{
				CardStack: []cards.Card{{Code: uint8(1)}},
			},
			args:    args{removeNum: 0},
			wantErr: false,
		},
		{
			name: "filled: move 1 OK",
			fields: fields{
				CardStack: []cards.Card{{Code: uint8(1)}},
			},
			args:    args{removeNum: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hCell := HomeCell{
				CardStack: tt.fields.CardStack,
				SuitCode:  tt.fields.SuitCode,
			}
			if err := hCell.RemoveEndSeq(tt.args.removeNum); (err != nil) != tt.wantErr {
				t.Errorf("RemoveEndSeq() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHomeCell_Place(t *testing.T) {
	type fields struct {
		CardStack []cards.Card
		SuitCode  uint8
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
				CardStack: []cards.Card{{Code: uint8(1)}},
			},
			args: args{
				seq: &[]cards.Card{{Code: uint8(2)}},
			},
			want: []cards.Card{
				{Code: uint8(1)},
				{Code: uint8(2)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hCell := HomeCell{
				CardStack: tt.fields.CardStack,
				SuitCode:  tt.fields.SuitCode,
			}

			hCell.Place(tt.args.seq)
			if !EqualStack(&(hCell.CardStack), &(tt.want)) {
				t.Errorf("Place() = %v, want %v", hCell.CardStack, tt.want)
			}
		})
	}
}
