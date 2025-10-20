package main

import (
	"fmt"
)

func main() {
	const threshold = 12307
	var num int
	fmt.Scan(&num)
	if num >= threshold {
		fmt.Println("The number is already above the threshold, no calculations are needed.")
		return
	}
	for num < threshold {
		switch {
		case num < 0:
			num *= -1
		case num%7 == 0:
			num *= 39
		case num%9 == 0:
			num = num*13 + 1
			continue
		default:
			num = (num + 2) * 3
		}
		if num%9 == 0 && num%13 == 0 {
			fmt.Println("service error")
			return
		} else {
			num++
		}
	}

	fmt.Printf("After complicated calculations the number become: %d. Congratulations!\n", num)
}
