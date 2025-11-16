//go:build amd64 || arm64 || mips64 || mips64le || ppc64 || ppc64le || riscv64 || s390x || wasm

package summator

import (
	"context"
	"sync"
	"sync/atomic"
)

func sumSquares(start, end, step int) int {
	var sum int64
	l, r, st := int64(start), int64(end), int64(step)
	nums := make(chan int64)
	wg := sync.WaitGroup{}
	ctx := context.Context(context.Background())
	ctx, cancel := context.WithCancel(ctx)
	wg.Add(workersCount)
	for range workersCount {
		go addSquare64(ctx, &wg, nums, &sum)
	}
	for i := l; i <= r; i += st {
		nums <- i
	}
	cancel()
	wg.Wait()
	return int(sum)
}

func addSquare64(ctx context.Context, wg *sync.WaitGroup, in <-chan int64, sum *int64) {
	for {
		select {
		case n := <-in:
			inc := n * n
			atomic.AddInt64(sum, inc)
		case <-ctx.Done():
			wg.Done()
			return
		}
	}
}
