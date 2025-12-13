//go:build arm64 || amd64

package summator

import (
	"testing"
)

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
			name:     "both zero",
			start:    0,
			end:      0,
			expected: 0,
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
			name:     "rearranged start, end",
			start:    5,
			end:      1,
			expected: 25 + 16 + 9 + 4 + 1, // 55
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

func TestSumSquares_StressTest(t *testing.T) {
	const n = 1000000
	tests := []struct {
		name     string
		start    int
		end      int
		expected int
	}{
		{
			name:     "positive range",
			start:    0,
			end:      n,
			expected: n * (n + 1) * (2*n + 1) / 6,
		},
		{
			name:     "negative range",
			start:    -n,
			end:      0,
			expected: n * (n + 1) * (2*n + 1) / 6,
		},
		{
			name:     "double range",
			start:    -1000000,
			end:      1000000,
			expected: 2 * n * (n + 1) * (2*n + 1) / 6,
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

func TestSumSquaresWithStep_BasicFunctionality(t *testing.T) {
	tests := []struct {
		name     string
		start    int
		end      int
		step     int
		expected int
	}{
		{
			name:     "single number",
			start:    5,
			end:      5,
			step:     1,
			expected: 25,
		},
		{
			name:     "both zero",
			start:    0,
			end:      0,
			step:     5,
			expected: 0,
		},
		{
			name:     "small range",
			start:    1,
			end:      3,
			step:     2,
			expected: 1 + 9, // 10
		},
		{
			name:     "range including zero",
			start:    0,
			end:      2,
			step:     2,
			expected: 0 + 4, // 4
		},
		{
			name:     "negative to positive",
			start:    -3,
			end:      3,
			step:     3,
			expected: 9 + 0 + 9, // 18
		},
		{
			name:     "all negatives",
			start:    -5,
			end:      -1,
			step:     2,
			expected: 25 + 9 + 1, // 35
		},
		{
			name:     "rearranged start, end",
			start:    5,
			end:      1,
			step:     2,
			expected: 25 + 9 + 1, // 35
		},
		{
			name:     "negative step",
			start:    2,
			end:      10,
			step:     -4,
			expected: 4 + 36 + 100, // 140
		},
		{
			name:     "task text example",
			start:    2,
			end:      10,
			step:     2,
			expected: 4 + 16 + 36 + 64 + 100, // 220
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SumSquaresWithStep(tt.start, tt.end, tt.step)
			if result != tt.expected {
				t.Errorf("SumSquares(%d, %d, %d) = %d, expected %d",
					tt.start, tt.end, tt.step, result, tt.expected)
			}
		})
	}
}

func TestSumSquaresWithStep_StressTest(t *testing.T) {
	const n = 1000000
	tests := []struct {
		name     string
		start    int
		end      int
		step     int
		expected int
	}{
		{
			name:     "even range",
			start:    0,
			end:      2 * n,
			step:     2,
			expected: 2 * n * (n + 1) * (2*n + 1) / 3,
		},
		{
			name:     "odd range",
			start:    1,
			end:      2*n - 1,
			step:     2,
			expected: n * (2*n - 1) * (2*n + 1) / 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SumSquaresWithStep(tt.start, tt.end, tt.step)
			if result != tt.expected {
				t.Errorf("SumSquaresWithStep(%d, %d, %d) = %d, expected %d",
					tt.start, tt.end, tt.step, result, tt.expected)
			}
		})
	}
}
