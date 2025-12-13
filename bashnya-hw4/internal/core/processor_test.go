package core

import (
	"testing"
)

func TestSeekUnique_BasicFunctionality(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "empty input",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "single line",
			input:    []string{"hello"},
			expected: []string{"hello"},
		},
		{
			name:     "basic duplicates",
			input:    []string{"a", "a", "b", "b", "b", "c"},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "no consecutive duplicates",
			input:    []string{"a", "b", "a", "b"},
			expected: []string{"a", "b", "a", "b"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SeekUnique(tt.input, &Options{})
			if !equalSlices(result, tt.expected) {
				t.Errorf("SeekUnique(%v) = %v, expected %v",
					tt.input, result, tt.expected)
			}
		})
	}
}

func TestSeekUnique_CountFlag(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "count basic duplicates",
			input:    []string{"a", "a", "b", "b", "b", "c"},
			expected: []string{"2 a", "3 b", "1 c"},
		},
		{
			name:     "count single line",
			input:    []string{"hello"},
			expected: []string{"1 hello"},
		},
		{
			name:     "count mixed duplicates",
			input:    []string{"a", "b", "a", "b", "c"},
			expected: []string{"1 a", "1 b", "1 a", "1 b", "1 c"},
		},
		{
			name:     "count empty lines",
			input:    []string{"", "", "a", "a", ""},
			expected: []string{"2 ", "2 a", "1 "},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := &Options{Count: true}
			result := SeekUnique(tt.input, opts)
			if !equalSlices(result, tt.expected) {
				t.Errorf("SeekUnique with -c: got %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestSeekUnique_RepeatedFlag(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "repeated basic duplicates",
			input:    []string{"a", "a", "b", "b", "b", "c"},
			expected: []string{"a", "b"},
		},
		{
			name:     "repeated single line",
			input:    []string{"hello"},
			expected: []string{},
		},
		{
			name:     "repeated no duplicates",
			input:    []string{"a", "b", "c"},
			expected: []string{},
		},
		{
			name:     "repeated mixed pattern",
			input:    []string{"a", "a", "b", "c", "c", "d"},
			expected: []string{"a", "c"},
		},
		{
			name:     "repeated multiple groups",
			input:    []string{"a", "a", "b", "a", "a"},
			expected: []string{"a", "a"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := &Options{Repeated: true}
			result := SeekUnique(tt.input, opts)
			if !equalSlices(result, tt.expected) {
				t.Errorf("SeekUnique with -d: got %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestSeekUnique_UniqueFlag(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "unique basic duplicates",
			input:    []string{"a", "a", "b", "b", "b", "c"},
			expected: []string{"c"},
		},
		{
			name:     "unique single line",
			input:    []string{"hello"},
			expected: []string{"hello"},
		},
		{
			name:     "unique all duplicates",
			input:    []string{"a", "a", "a"},
			expected: []string{},
		},
		{
			name:     "unique mixed pattern",
			input:    []string{"a", "a", "b", "c", "c", "d"},
			expected: []string{"b", "d"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := &Options{Unique: true}
			result := SeekUnique(tt.input, opts)
			if !equalSlices(result, tt.expected) {
				t.Errorf("SeekUnique with -u: got %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestSeekUnique_IgnoreCaseFlag(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
		opts     *Options
	}{
		{
			name:     "ignore case basic",
			input:    []string{"Hello", "HELLO", "hello", "World", "world"},
			expected: []string{"Hello", "World"},
			opts:     &Options{IgnoreCase: true},
		},
		{
			name:     "ignore case with count",
			input:    []string{"A", "a", "B", "b", "b"},
			expected: []string{"2 A", "3 B"},
			opts:     &Options{IgnoreCase: true, Count: true},
		},
		{
			name:     "ignore case with repeated",
			input:    []string{"test", "TEST", "hello", "Hello"},
			expected: []string{"test", "hello"},
			opts:     &Options{IgnoreCase: true, Repeated: true},
		},
		{
			name:     "ignore case with unique",
			input:    []string{"a", "A", "b", "B", "c"},
			expected: []string{"c"},
			opts:     &Options{IgnoreCase: true, Unique: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SeekUnique(tt.input, tt.opts)
			if !equalSlices(result, tt.expected) {
				t.Errorf("SeekUnique with -i: got %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestSeekUnique_SkipFieldsFlag(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
		opts     *Options
	}{
		{
			name:     "skip fields basic",
			input:    []string{"apple fruit", "banana fruit", "apple computer", "banana computer"},
			expected: []string{"apple fruit", "apple computer"},
			opts:     &Options{SkipFields: 1},
		},
		{
			name:     "skip fields with count",
			input:    []string{"1 apple", "2 apple", "3 banana", "4 banana"},
			expected: []string{"2 1 apple", "2 3 banana"},
			opts:     &Options{SkipFields: 1, Count: true},
		},
		{
			name:     "skip multiple fields",
			input:    []string{"a b c", "a b d", "x y c", "x y d"},
			expected: []string{"a b c", "a b d", "x y c", "x y d"},
			opts:     &Options{SkipFields: 2},
		},
		{
			name:     "skip fields more than available",
			input:    []string{"simple", "simple", "different"},
			expected: []string{"simple"},
			opts:     &Options{SkipFields: 10},
		},
		{
			name:     "skip fields with repeated",
			input:    []string{"a x", "a y", "b x", "b y"},
			expected: []string{},
			opts:     &Options{SkipFields: 1, Repeated: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SeekUnique(tt.input, tt.opts)
			if !equalSlices(result, tt.expected) {
				t.Errorf("SeekUnique with -f: got %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestSeekUnique_SkipCharsFlag(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		opts     *Options
		expected []string
	}{
		{
			name:     "skip chars basic",
			input:    []string{"abcbashnya", "cdfbashnya", "xyzbashnya", "123bashnya"},
			expected: []string{"abcbashnya"},
			opts:     &Options{SkipChars: 3},
		},
		{
			name:     "skip chars with count",
			input:    []string{"prefix1", "suffix1", "golang1", "something"},
			expected: []string{"3 prefix1", "1 something"},
			opts:     &Options{SkipChars: 6, Count: true},
		},
		{
			name:     "skip more chars than string length",
			input:    []string{"hi", "hi", "hello"},
			expected: []string{"hi"},
			opts:     &Options{SkipChars: 10},
		},
		{
			name:     "skip chars with ignore case",
			input:    []string{"ABCDATA", "abcdata", "XYZdata"},
			expected: []string{"ABCDATA"},
			opts:     &Options{SkipChars: 3, IgnoreCase: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SeekUnique(tt.input, tt.opts)
			if !equalSlices(result, tt.expected) {
				t.Errorf("SeekUnique with -s: got %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestSeekUnique_CombinedFlags(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		opts     *Options
		expected []string
	}{
		{
			name:     "skip fields and chars",
			input:    []string{"aa bb cc", "dd bb cc", "xx yy cc", "dd yy cc"},
			opts:     &Options{SkipFields: 1, SkipChars: 1},
			expected: []string{"aa bb cc", "xx yy cc"},
		},
		{
			name:     "ignore case with skip fields",
			input:    []string{"A B C", "a B C", "X Y Z", "x Y Z"},
			opts:     &Options{IgnoreCase: true, SkipFields: 1},
			expected: []string{"A B C", "X Y Z"},
		},
		{
			name:     "multiple skips with tabs and spaces",
			input:    []string{"  apple  red", "orange  red", "  grape red", "banana yellow"},
			opts:     &Options{SkipFields: 1, SkipChars: 2},
			expected: []string{"  apple  red", "  grape red", "banana yellow"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SeekUnique(tt.input, tt.opts)
			if !equalSlices(result, tt.expected) {
				t.Errorf("SeekUnique with combined flags: got %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestSeekUnique_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		opts     *Options
		expected []string
	}{
		{
			name:     "all empty lines",
			input:    []string{"", "", ""},
			opts:     &Options{},
			expected: []string{""},
		},
		{
			name:     "empty lines with count",
			input:    []string{"", "", ""},
			opts:     &Options{Count: true},
			expected: []string{"3 "},
		},
		{
			name:     "mixed empty and non-empty",
			input:    []string{"", "a", "", "a", ""},
			opts:     &Options{},
			expected: []string{"", "a", "", "a", ""},
		},
		{
			name:     "whitespace variations",
			input:    []string{"  hello", "  hello", "hello", "hello  "},
			opts:     &Options{},
			expected: []string{"  hello", "hello", "hello  "},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SeekUnique(tt.input, tt.opts)
			if !equalSlices(result, tt.expected) {
				t.Errorf("Edge case %s: got %v, expected %v", tt.name, result, tt.expected)
			}
		})
	}
}

func TestSeekUnique_MutuallyExclusiveFlags(t *testing.T) {
	// Test that the function handles flag validation properly
	// This would typically be tested in NewUniqProcessor, but let's ensure SeekUnique behaves reasonably

	input := []string{"a", "a", "b", "b", "c"}

	t.Run("count takes precedence when multiple set", func(t *testing.T) {
		opts := &Options{Count: true, Repeated: true, Unique: true}
		result := SeekUnique(input, opts)
		// Should behave as if only Count was set
		expected := []string{"2 a", "2 b", "1 c"}
		if !equalSlices(result, expected) {
			t.Errorf("With multiple flags: got %v, expected %v", result, expected)
		}
	})
}

// Helper function to compare string slices
func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
