package main

import (
	"testing"
)

func TestLevelGlyph(t *testing.T) {
	tests := []struct {
		capacity int
		want     string
	}{
		{100, "█"},
		{87, "▇"},
		{74, "▆"},
		{62, "▅"},
		{49, "▄"},
		{37, "▃"},
		{24, "▂"},
		{0, "▁"},
	}
	for _, test := range tests {
		if got := LevelGlyph(test.capacity); got != test.want {
			t.Errorf("LevelGlyph(%d) = %v", test.capacity, got)
		}
	}
}

func TestLevelIndex(t *testing.T) {
	var tests = []struct {
		capacity int
		want     int
	}{
		{101, 7},
		{100, 7},
		{99, 7},
		{88, 7},

		{87, 6},
		{75, 6},

		{74, 5},
		{63, 5},

		{62, 4},
		{50, 4},

		{49, 3},
		{38, 3},

		{37, 2},
		{25, 2},

		{24, 1},
		{13, 1},

		{12, 0},
		{0, 0},
		{-1, 0},
	}
	for _, test := range tests {
		if got := levelIndex(test.capacity); got != test.want {
			t.Errorf("LevelIndex(%d) = %v", test.capacity, got)
		}
	}
}
