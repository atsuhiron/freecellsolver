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
			name:   "GetSuitCode 1",
			fields: fields{uint8(1)}, // ♠A
			want:   uint8(0),
		},
		{
			name:   "GetSuitCode 2",
			fields: fields{uint8(2)}, // ♠2
			want:   uint8(0),
		},
		{
			name:   "GetSuitCode 3",
			fields: fields{uint8(33)}, // ♥A
			want:   uint8(1),
		},
		{
			name:   "GetSuitCode 4",
			fields: fields{uint8(74)}, // ♣10
			want:   uint8(2),
		},
		{
			name:   "GetSuitCode 5",
			fields: fields{uint8(109)}, // ♦K
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
			name:   "IsBlack Spade",
			fields: fields{uint8(1)}, // ♠A
			want:   true,
		},
		{
			name:   "IsBlack Heart",
			fields: fields{uint8(33)}, // ♥A
			want:   false,
		},
		{
			name:   "IsBlack Diamonds",
			fields: fields{uint8(74)}, // ♣10
			want:   true,
		},
		{
			name:   "IsBlack Clover",
			fields: fields{uint8(109)}, // ♦K
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
			name:   "ToReadableCard 1",
			fields: fields{uint8(1)}, // ♠A
			want:   ReadableCard{"♠", "A"},
		},
		{
			name:   "ToReadableCard 2",
			fields: fields{uint8(2)}, // ♠2
			want:   ReadableCard{"♠", "2"},
		},
		{
			name:   "ToReadableCard 3",
			fields: fields{uint8(33)}, // ♥A
			want:   ReadableCard{"♥", "A"},
		},
		{
			name:   "ToReadableCard 4",
			fields: fields{uint8(74)}, // ♣10
			want:   ReadableCard{"♣", "10"},
		},
		{
			name:   "ToReadableCard 5",
			fields: fields{uint8(109)}, // ♦K
			want:   ReadableCard{"♦", "K"},
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
