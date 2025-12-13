package bits

import (
	"testing"
)

func TestSet32(t *testing.T) {
	tests := []struct {
		name     string
		n        int32
		pos      int
		expected int32
	}{
		{
			name:     "basic",
			n:        32,
			pos:      0,
			expected: 33,
		},
		{
			name:     "no change",
			n:        32,
			pos:      5,
			expected: 32,
		},
		{
			name:     "big result",
			n:        32,
			pos:      30,
			expected: 1<<30 + 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if exp := Set32(tt.n, tt.pos); exp != tt.expected {
				t.Errorf("Set32(%d, %d) = %d. Expected: %d\n", tt.n, tt.pos, exp, tt.expected)
			}
		})
	}
}

func TestReset32(t *testing.T) {
	tests := []struct {
		name     string
		n        int32
		pos      int
		expected int32
	}{
		{
			name:     "basic",
			n:        33,
			pos:      0,
			expected: 32,
		},
		{
			name:     "no change",
			n:        32,
			pos:      2,
			expected: 32,
		},
		{
			name:     "zero result",
			n:        32,
			pos:      5,
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if exp := Reset32(tt.n, tt.pos); exp != tt.expected {
				t.Errorf("Set32(%d, %d) = %d. Expected: %d\n", tt.n, tt.pos, exp, tt.expected)
			}
		})
	}
}

func TestFlip32(t *testing.T) {
	tests := []struct {
		name     string
		n        int32
		pos      int
		expected int32
	}{
		{
			name:     "basic",
			n:        32,
			pos:      0,
			expected: 33,
		},
		{
			name:     "zero result",
			n:        32,
			pos:      5,
			expected: 0,
		},
		{
			name:     "big result",
			n:        32,
			pos:      30,
			expected: 1<<30 + 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if exp := Flip32(tt.n, tt.pos); exp != tt.expected {
				t.Errorf("Set32(%d, %d) = %d. Expected: %d\n", tt.n, tt.pos, exp, tt.expected)
			}
		})
	}
}
