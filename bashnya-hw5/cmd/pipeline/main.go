package main

import (
	"math"
	"os"

	"github.com/vitalykov/bashnya-backend-basic/bashnya-hw5/pipeline"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	pipeline.WriteTransformed(nums, os.Stdout, func(a int) int {
		return a * a
	})
	pipeline.WriteTransformed(nums, os.Stdout, func(a int) int {
		return a + 1
	})
	pipeline.WriteTransformed(nums, os.Stdout, func(a int) int {
		return int(math.Pow10(a))
	})
	pipeline.WriteTransformed(nums, os.Stdout, func(a int) int {
		return int(math.Pow(float64(a), float64(a)))
	})
}
