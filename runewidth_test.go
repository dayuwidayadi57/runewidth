package runewidth

import "testing"

func TestGetVisualWidth(t *testing.T) {
	tests := []struct {
		r    rune
		want int
	}{
		{'A', 1},
		{'\U0001F680', 2},
		{'\u2705', 2},
		{'\u4e2d', 2},
	}
	for _, tt := range tests {
		if got := GetVisualWidth(tt.r); got != tt.want {
			t.Errorf("GetVisualWidth(%U) = %d, want %d", tt.r, got, tt.want)
		}
	}
}

func TestStringVisualWidth(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"Hello", 5},
		{"Hello \U0001F680", 8},
	}
	for _, tt := range tests {
		if got := StringVisualWidth(tt.s); got != tt.want {
			t.Errorf("StringVisualWidth(%q) = %d, want %d", tt.s, got, tt.want)
		}
	}
}

