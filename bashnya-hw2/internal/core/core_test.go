package core

import "testing"

// Could break if threshold changed
func TestServiceErrors(t *testing.T) {
	numbers := [...]int{-117, 12306}
	for _, num := range numbers {
		res, err := MakeCalculations(num)
		if err == nil {
			t.Errorf("MakeCalculations(%d) = %d. Expected error: %s", num, res, errorMsg)
		}
	}
}

// Could break if threshold changed
func TestCalculations(t *testing.T) {
	numbers := map[int]int{
		-threshold: threshold + 1,
		threshold:  threshold,
		12299:      479662,
		12303:      159940,
		12305:      36922,
	}
	for num, expected := range numbers {
		res, err := MakeCalculations(num)
		if err != nil {
			t.Errorf("MakeCalculations(%d) = %s. Expected: %d", num, err.Error(), expected)
		}
		if res != expected {
			t.Errorf("MakeCalculations(%d) = %d. Expected: %d", num, res, expected)
		}
	}
}
