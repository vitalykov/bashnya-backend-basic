package summator

import (
	"runtime"
)

var workersCount = runtime.NumCPU()

func SumSquaresWithStep(start, end, step int) int {
	// different signs have end - start and step
	if (end-start)*step < 0 || step == 0 {
		return -1
	}
	return sumSquares(start, end, step)
}

func SumSquares(start, end int) int {
	if end < start {
		start, end = end, start
	}
	return sumSquares(start, end, 1)
}
