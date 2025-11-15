//go:build 386 || arm || mips

package summator

import (
	"context"
	"sync"
	"sync/atomic"
)

func sumSquares(start, end, step int) int {
	var sum int32
	l, r, st := int32(start), int32(end), int32(step)
	nums := make(chan int32)
	wg := sync.WaitGroup{}
	ctx := context.Context(context.Background())
	ctx, cancel := context.WithCancel(ctx)
	wg.Add(workersCount)
	for range workersCount {
		go addSquare32(ctx, &wg, nums, &sum)
	}
	for i := l; i <= r; i += st {
		nums <- i
	}
	cancel()
	wg.Wait()
	return int(sum)
}

func addSquare32(ctx context.Context, wg *sync.WaitGroup, in <-chan int32, sum *int32) {
	for {
		select {
		case n := <-in:
			inc := n * n
			atomic.AddInt32(sum, inc)
		case <-ctx.Done():
			wg.Done()
			return
		}
	}
}
