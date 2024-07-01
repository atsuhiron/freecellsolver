package cards

import "testing"

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
			name:   "Get Suit Code 1",
			fields: fields{uint8(1)}, // ♠A
			want:   uint8(0),
		},
		{
			name:   "Get Suit Code 2",
			fields: fields{uint8(2)}, // ♠2
			want:   uint8(0),
		},
		{
			name:   "Get Suit Code 3",
			fields: fields{uint8(33)}, // ♥A
			want:   uint8(1),
		},
		{
			name:   "Get Suit Code 4",
			fields: fields{uint8(69)}, // ♣10
			want:   uint8(2),
		},
		{
			name:   "Get Suit Code 5",
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
