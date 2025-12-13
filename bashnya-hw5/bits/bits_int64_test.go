package bits

import (
	"testing"
)

func TestSet64(t *testing.T) {
	tests := []struct {
		name     string
		n        int64
		pos      int
		expected int64
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
			pos:      62,
			expected: 1<<62 + 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if exp := Set64(tt.n, tt.pos); exp != tt.expected {
				t.Errorf("Set64(%d, %d) = %d. Expected: %d\n", tt.n, tt.pos, exp, tt.expected)
			}
		})
	}
}

func TestReset64(t *testing.T) {
	tests := []struct {
		name     string
		n        int64
		pos      int
		expected int64
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
			if exp := Reset64(tt.n, tt.pos); exp != tt.expected {
				t.Errorf("Set64(%d, %d) = %d. Expected: %d\n", tt.n, tt.pos, exp, tt.expected)
			}
		})
	}
}

func TestFlip64(t *testing.T) {
	tests := []struct {
		name     string
		n        int64
		pos      int
		expected int64
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
			pos:      62,
			expected: 1<<62 + 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if exp := Flip64(tt.n, tt.pos); exp != tt.expected {
				t.Errorf("Set64(%d, %d) = %d. Expected: %d\n", tt.n, tt.pos, exp, tt.expected)
			}
		})
	}
}
