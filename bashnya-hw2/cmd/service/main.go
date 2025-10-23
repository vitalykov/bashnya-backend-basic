package main

import (
	"fmt"

	"github.com/vitalykov/bashnya-backend-basic/bashnya-hw2/internal/core"
	"github.com/vitalykov/bashnya-backend-basic/bashnya-hw2/numbers"
)

func main() {
	var num int
	fmt.Scan(&num)
	result, err := core.MakeCalculations(num)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if num != result {
		fmt.Printf("After complicated calculations the number become: %d. Congratulations!\n", result)
	} else {
		fmt.Println("The number is already above the threshold, no calculations are needed.")
	}
	fmt.Printf("А теперь по-русски: %s\n", numbers.IntToWords(result))
}
