package summator

import (
	"testing"
)

// TestSumSquares_BasicFunctionality tests normal cases
func TestSumSquares_BasicFunctionality(t *testing.T) {
	tests := []struct {
		name     string
		start    int
		end      int
		expected int
	}{
		{
			name:     "single number",
			start:    5,
			end:      5,
			expected: 25,
		},
		{
			name:     "small range",
			start:    1,
			end:      3,
			expected: 1 + 4 + 9, // 14
		},
		{
			name:     "range including zero",
			start:    0,
			end:      2,
			expected: 0 + 1 + 4, // 5
		},
		{
			name:     "negative to positive",
			start:    -2,
			end:      2,
			expected: 4 + 1 + 0 + 1 + 4, // 10
		},
		{
			name:     "all negatives",
			start:    -3,
			end:      -1,
			expected: 9 + 4 + 1, // 14
		},
		{
			name:     "large numbers but safe",
			start:    10,
			end:      12,
			expected: 100 + 121 + 144, // 365
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SumSquares(tt.start, tt.end)
			if result != tt.expected {
				t.Errorf("SumSquares(%d, %d) = %d, expected %d",
					tt.start, tt.end, result, tt.expected)
			}
		})
	}
}

func TestSumSquares_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		start    int
		end      int
		expected int
	}{
		{
			name:     "both zero",
			start:    0,
			end:      0,
			expected: 0,
		},
		{
			name:     "range of length 1 at extreme negative",
			start:    -10000,
			end:      -10000,
			expected: 10000 * 10000,
		},
		{
			name:     "range of length 1 at extreme positive",
			start:    10000,
			end:      10000,
			expected: 10000 * 10000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SumSquares(tt.start, tt.end)

			if result != tt.expected {
				t.Errorf("SumSquares(%d, %d) = %d, expected %d",
					tt.start, tt.end, result, tt.expected)
			}
		})
	}
}
