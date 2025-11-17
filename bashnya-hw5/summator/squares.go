package summator

import (
	"runtime"
	"sync"
)

var workersCount = runtime.GOMAXPROCS(-1)

func SumSquaresWithStep(start, end, step int) int {
	if end < start {
		start, end = end, start
	}
	if step < 0 {
		step = -step
	}
	return sumSquares(start, end, step)
}

func SumSquares(start, end int) int {
	if end < start {
		start, end = end, start
	}
	return sumSquares(start, end, 1)
}

func sumSquares(start, end, step int) int {
	totalSteps := (end - start) / step
	intervalsCount := min(totalSteps+1, workersCount)
	size := totalSteps / intervalsCount
	partEnd := start + size*step
	sumParts := make(chan int, intervalsCount)
	var wg sync.WaitGroup
	wg.Add(intervalsCount)
	for i := 0; i < intervalsCount-1; i++ {
		go partSumSquares(&wg, sumParts, start, partEnd, step)
		start = partEnd + step
		partEnd = start + size*step
	}
	go partSumSquares(&wg, sumParts, start, end, step)
	wg.Wait()
	close(sumParts)
	var sum int
	for part := range sumParts {
		sum += part
	}
	return sum
}

func partSumSquares(wg *sync.WaitGroup, out chan<- int, start, end, step int) {
	var sum int
	for i := start; i <= end; i += step {
		sum += i * i
	}
	out <- sum
	wg.Done()
}
