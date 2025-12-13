package pipeline

import (
	"fmt"
	"io"
)

func WriteTransformed(nums []int, w io.Writer, op func(int) int) {
	in := make(chan int)
	out := make(chan int)

	go func() {
		for _, n := range nums {
			in <- n
		}
		close(in)
	}()

	go func() {
		for n := range in {
			out <- op(n)
		}
		close(out)
	}()

	for n := range out {
		fmt.Fprintf(w, "%d ", n)
	}
	fmt.Fprintln(w)
}
