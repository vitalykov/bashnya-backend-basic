package printer

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"sync"
	"time"
)

type Printer struct {
	inChan       chan any
	out          io.Writer
	workersCount int
	wg           *sync.WaitGroup
	ctx          context.Context
	cancel       context.CancelFunc
}

func NewPrinter(in chan any, out io.Writer, workersCount int) *Printer {
	ctx := context.Context(context.Background())
	ctx, cancel := context.WithCancel(ctx)
	p := Printer{
		inChan:       in,
		out:          out,
		workersCount: workersCount,
		wg:           &sync.WaitGroup{},
		ctx:          ctx,
		cancel:       cancel,
	}
	return &p
}

func (p *Printer) print() {
	for {
		select {
		case val := <-p.inChan:
			fmt.Fprintln(p.out, val)
			time.Sleep(time.Duration(int(time.Millisecond) * rand.Intn(1000)))
		case <-p.ctx.Done():
			p.wg.Done()
			return
		}
	}
}

func (p *Printer) Run() {
	p.wg.Add(p.workersCount)
	for range p.workersCount {
		go p.print()
	}
}

func (p *Printer) Shutdown() {
	p.cancel()
	p.wg.Wait()
	close(p.inChan)
}
