package summator

import (
	"runtime"
)

var workersCount = runtime.NumCPU()

func SumSquares(start, end int) int {
	if end < start {
		start, end = end, start
	}
	return sumSquares(start, end)
}
