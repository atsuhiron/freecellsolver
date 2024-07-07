package cards

import (
	"reflect"
	"testing"
)

func TestReadableCard_ToCard(t *testing.T) {
	type fields struct {
		Suit string
		Num  string
	}
	tests := []struct {
		name    string
		fields  fields
		want    Card
		wantErr bool
	}{
		{
			name:    "♠A",
			fields:  fields{"♠", "A"},
			want:    Card{uint8(1)},
			wantErr: false,
		},
		{
			name:    "SA",
			fields:  fields{"S", "A"},
			want:    Card{uint8(1)},
			wantErr: false,
		},
		{
			name:    "♠1",
			fields:  fields{"♠", "1"},
			want:    Card{uint8(1)},
			wantErr: false,
		},
		{
			name:    "S1",
			fields:  fields{"S", "1"},
			want:    Card{uint8(1)},
			wantErr: false,
		},
		{
			name:    "♠2",
			fields:  fields{"♠", "2"},
			want:    Card{uint8(2)},
			wantErr: false,
		},
		{
			name:    "S2",
			fields:  fields{"S", "2"},
			want:    Card{uint8(2)},
			wantErr: false,
		},
		{
			name:    "♠10",
			fields:  fields{"♠", "10"},
			want:    Card{uint8(10)},
			wantErr: false,
		},
		{
			name:    "S10",
			fields:  fields{"S", "10"},
			want:    Card{uint8(10)},
			wantErr: false,
		},
		{
			name:    "♠T",
			fields:  fields{"♠", "T"},
			want:    Card{uint8(10)},
			wantErr: false,
		},
		{
			name:    "ST",
			fields:  fields{"S", "T"},
			want:    Card{uint8(10)},
			wantErr: false,
		},
		{
			name:    "♠11",
			fields:  fields{"♠", "11"},
			want:    Card{uint8(11)},
			wantErr: false,
		},
		{
			name:    "S11",
			fields:  fields{"S", "11"},
			want:    Card{uint8(11)},
			wantErr: false,
		},
		{
			name:    "♠J",
			fields:  fields{"♠", "J"},
			want:    Card{uint8(11)},
			wantErr: false,
		},
		{
			name:    "SJ",
			fields:  fields{"S", "J"},
			want:    Card{uint8(11)},
			wantErr: false,
		},
		{
			name:    "♠12",
			fields:  fields{"♠", "12"},
			want:    Card{uint8(12)},
			wantErr: false,
		},
		{
			name:    "S12",
			fields:  fields{"S", "12"},
			want:    Card{uint8(12)},
			wantErr: false,
		},
		{
			name:    "♠Q",
			fields:  fields{"♠", "Q"},
			want:    Card{uint8(12)},
			wantErr: false,
		},
		{
			name:    "SQ",
			fields:  fields{"S", "Q"},
			want:    Card{uint8(12)},
			wantErr: false,
		},
		{
			name:    "♠13",
			fields:  fields{"♠", "13"},
			want:    Card{uint8(13)},
			wantErr: false,
		},
		{
			name:    "S13",
			fields:  fields{"S", "13"},
			want:    Card{uint8(13)},
			wantErr: false,
		},
		{
			name:    "♠K",
			fields:  fields{"♠", "K"},
			want:    Card{uint8(13)},
			wantErr: false,
		},
		{
			name:    "SK",
			fields:  fields{"S", "K"},
			want:    Card{uint8(13)},
			wantErr: false,
		},
		{
			name:    "♥5",
			fields:  fields{"♥", "5"},
			want:    Card{uint8(21)},
			wantErr: false,
		},
		{
			name:    "H5",
			fields:  fields{"H", "5"},
			want:    Card{uint8(21)},
			wantErr: false,
		},
		{
			name:    "♣10",
			fields:  fields{"♣", "10"},
			want:    Card{uint8(42)},
			wantErr: false,
		},
		{
			name:    "C10",
			fields:  fields{"C", "10"},
			want:    Card{uint8(42)},
			wantErr: false,
		},
		{
			name:    "♦K",
			fields:  fields{"♦", "K"},
			want:    Card{uint8(61)},
			wantErr: false,
		},
		{
			name:    "DK",
			fields:  fields{"D", "K"},
			want:    Card{uint8(61)},
			wantErr: false,
		},
		{
			name:    "Error ♠X",
			fields:  fields{"♠", "X"},
			want:    Card{uint8(255)},
			wantErr: true,
		},
		{
			name:    "Error AB",
			fields:  fields{"A", "B"},
			want:    Card{uint8(255)},
			wantErr: true,
		},
		{
			name:    "Error ♠99",
			fields:  fields{"♠", "99"},
			want:    Card{uint8(255)},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rCard := ReadableCard{
				Suit: tt.fields.Suit,
				Num:  tt.fields.Num,
			}
			got, err := rCard.ToCard()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToCard() got = %v, want %v", got, tt.want)
			}
		})
	}
}
