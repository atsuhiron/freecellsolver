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
			name:    "ToCard ♠A",
			fields:  fields{"♠", "A"},
			want:    Card{uint8(1)},
			wantErr: false,
		},
		{
			name:    "ToCard SA",
			fields:  fields{"S", "A"},
			want:    Card{uint8(1)},
			wantErr: false,
		},
		{
			name:    "ToCard ♠1",
			fields:  fields{"♠", "1"},
			want:    Card{uint8(1)},
			wantErr: false,
		},
		{
			name:    "ToCard S1",
			fields:  fields{"S", "1"},
			want:    Card{uint8(1)},
			wantErr: false,
		},
		{
			name:    "ToCard ♠2",
			fields:  fields{"♠", "2"},
			want:    Card{uint8(2)},
			wantErr: false,
		},
		{
			name:    "ToCard S2",
			fields:  fields{"S", "2"},
			want:    Card{uint8(2)},
			wantErr: false,
		},
		{
			name:    "ToCard ♠10",
			fields:  fields{"♠", "10"},
			want:    Card{uint8(10)},
			wantErr: false,
		},
		{
			name:    "ToCard S10",
			fields:  fields{"S", "10"},
			want:    Card{uint8(10)},
			wantErr: false,
		},
		{
			name:    "ToCard ♠T",
			fields:  fields{"♠", "T"},
			want:    Card{uint8(10)},
			wantErr: false,
		},
		{
			name:    "ToCard ST",
			fields:  fields{"S", "T"},
			want:    Card{uint8(10)},
			wantErr: false,
		},
		{
			name:    "ToCard ♠11",
			fields:  fields{"♠", "11"},
			want:    Card{uint8(11)},
			wantErr: false,
		},
		{
			name:    "ToCard S11",
			fields:  fields{"S", "11"},
			want:    Card{uint8(11)},
			wantErr: false,
		},
		{
			name:    "ToCard ♠J",
			fields:  fields{"♠", "J"},
			want:    Card{uint8(11)},
			wantErr: false,
		},
		{
			name:    "ToCard SJ",
			fields:  fields{"S", "J"},
			want:    Card{uint8(11)},
			wantErr: false,
		},
		{
			name:    "ToCard ♠12",
			fields:  fields{"♠", "12"},
			want:    Card{uint8(12)},
			wantErr: false,
		},
		{
			name:    "ToCard S12",
			fields:  fields{"S", "12"},
			want:    Card{uint8(12)},
			wantErr: false,
		},
		{
			name:    "ToCard ♠Q",
			fields:  fields{"♠", "Q"},
			want:    Card{uint8(12)},
			wantErr: false,
		},
		{
			name:    "ToCard SQ",
			fields:  fields{"S", "Q"},
			want:    Card{uint8(12)},
			wantErr: false,
		},
		{
			name:    "ToCard ♠13",
			fields:  fields{"♠", "13"},
			want:    Card{uint8(13)},
			wantErr: false,
		},
		{
			name:    "ToCard S13",
			fields:  fields{"S", "13"},
			want:    Card{uint8(13)},
			wantErr: false,
		},
		{
			name:    "ToCard ♠K",
			fields:  fields{"♠", "K"},
			want:    Card{uint8(13)},
			wantErr: false,
		},
		{
			name:    "ToCard SK",
			fields:  fields{"S", "K"},
			want:    Card{uint8(13)},
			wantErr: false,
		},
		{
			name:    "ToCard ♥5",
			fields:  fields{"♥", "5"},
			want:    Card{uint8(21)},
			wantErr: false,
		},
		{
			name:    "ToCard H5",
			fields:  fields{"H", "5"},
			want:    Card{uint8(21)},
			wantErr: false,
		},
		{
			name:    "ToCard ♣10",
			fields:  fields{"♣", "10"},
			want:    Card{uint8(42)},
			wantErr: false,
		},
		{
			name:    "ToCard C10",
			fields:  fields{"C", "10"},
			want:    Card{uint8(42)},
			wantErr: false,
		},
		{
			name:    "ToCard ♦K",
			fields:  fields{"♦", "K"},
			want:    Card{uint8(61)},
			wantErr: false,
		},
		{
			name:    "ToCard DK",
			fields:  fields{"D", "K"},
			want:    Card{uint8(61)},
			wantErr: false,
		},
		{
			name:    "ToCard Error ♠X",
			fields:  fields{"♠", "X"},
			want:    Card{uint8(255)},
			wantErr: true,
		},
		{
			name:    "ToCard Error AB",
			fields:  fields{"A", "B"},
			want:    Card{uint8(255)},
			wantErr: true,
		},
		{
			name:    "ToCard Error ♠99",
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
