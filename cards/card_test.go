package cards

import (
	"reflect"
	"testing"
)

func TestCard_GetSuitCode(t *testing.T) {
	type fields struct {
		Code uint8
	}
	tests := []struct {
		name   string
		fields fields
		want   uint8
	}{
		{
			name:   "♠A",
			fields: fields{uint8(1)}, // ♠A
			want:   uint8(0),
		},
		{
			name:   "♠2",
			fields: fields{uint8(2)}, // ♠2
			want:   uint8(0),
		},
		{
			name:   "♥A",
			fields: fields{uint8(17)}, // ♥A
			want:   uint8(1),
		},
		{
			name:   "♣10",
			fields: fields{uint8(42)}, // ♣10
			want:   uint8(2),
		},
		{
			name:   "♦K",
			fields: fields{uint8(61)}, // ♦K
			want:   uint8(3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card := Card{
				Code: tt.fields.Code,
			}
			if got := card.GetSuitCode(); got != tt.want {
				t.Errorf("GetSuitCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_IsBlack(t *testing.T) {
	type fields struct {
		Code uint8
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "Spade",
			fields: fields{uint8(1)}, // ♠A
			want:   true,
		},
		{
			name:   "Heart",
			fields: fields{uint8(17)}, // ♥A
			want:   false,
		},
		{
			name:   "Clover",
			fields: fields{uint8(42)}, // ♣10
			want:   true,
		},
		{
			name:   "Diamonds",
			fields: fields{uint8(61)}, // ♦K
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card := Card{
				Code: tt.fields.Code,
			}
			if got := card.IsBlack(); got != tt.want {
				t.Errorf("IsBlack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_ToReadableCard(t *testing.T) {
	type fields struct {
		Code uint8
	}
	tests := []struct {
		name    string
		fields  fields
		want    ReadableCard
		wantErr bool
	}{
		{
			name:    "♠A",
			fields:  fields{uint8(1)}, // ♠A
			want:    ReadableCard{"♠", "A"},
			wantErr: false,
		},
		{
			name:    "♠2",
			fields:  fields{uint8(2)}, // ♠2
			want:    ReadableCard{"♠", "2"},
			wantErr: false,
		},
		{
			name:    "♥A",
			fields:  fields{uint8(17)}, // ♥A
			want:    ReadableCard{"♥", "A"},
			wantErr: false,
		},
		{
			name:    "♣10",
			fields:  fields{uint8(42)}, // ♣10
			want:    ReadableCard{"♣", "10"},
			wantErr: false,
		},
		{
			name:    "♦K",
			fields:  fields{uint8(61)}, // ♦K
			want:    ReadableCard{"♦", "K"},
			wantErr: false,
		},
		{
			name:    "♦Q",
			fields:  fields{uint8(60)}, // ♦K
			want:    ReadableCard{"♦", "Q"},
			wantErr: false,
		},
		{
			name:    "♦J",
			fields:  fields{uint8(59)}, // ♦K
			want:    ReadableCard{"♦", "J"},
			wantErr: false,
		},
		{
			name:    "error: out of suit range",
			fields:  fields{uint8(81)}, // out of range
			want:    ReadableCard{},
			wantErr: true,
		},
		{
			name:    "error: out of number range",
			fields:  fields{uint8(63)}, // out of range
			want:    ReadableCard{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card := Card{
				Code: tt.fields.Code,
			}
			got, err := card.ToReadableCard()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToReadableCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToReadableCard() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateCard(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    Card
		wantErr bool
	}{
		{
			name:    "Normal",
			args:    args{"S1"},
			want:    Card{uint8(1)},
			wantErr: false,
		},
		{
			name:    "Normal",
			args:    args{"S13"},
			want:    Card{uint8(13)},
			wantErr: false,
		},
		{
			name:    "Error: empty string",
			args:    args{""},
			want:    Card{uint8(255)},
			wantErr: true,
		},
		{
			name:    "Error: non-suit string",
			args:    args{"a"},
			want:    Card{uint8(255)},
			wantErr: true,
		},
		{
			name:    "Error: out of range number",
			args:    args{"C99"},
			want:    Card{uint8(255)},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateCard(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateCard() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_GetNumCode(t *testing.T) {
	type fields struct {
		Code uint8
	}
	tests := []struct {
		name   string
		fields fields
		want   uint8
	}{
		{
			name:   "♠A",
			fields: fields{uint8(1)}, // ♠A
			want:   uint8(1),
		},
		{
			name:   "♠2",
			fields: fields{uint8(2)}, // ♠2
			want:   uint8(2),
		},
		{
			name:   "♥A",
			fields: fields{uint8(17)}, // ♥A
			want:   uint8(1),
		},
		{
			name:   "♣10",
			fields: fields{uint8(42)}, // ♣10
			want:   uint8(10),
		},
		{
			name:   "♦K",
			fields: fields{uint8(61)}, // ♦K
			want:   uint8(13),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card := Card{
				Code: tt.fields.Code,
			}
			if got := card.GetNumCode(); got != tt.want {
				t.Errorf("GetNumCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
