package core

import "errors"

const (
	threshold = 12307
	errorMsg  = "service error"
)

func MakeCalculations(num int) (int, error) {
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
			return 0, errors.New(errorMsg)
		} else {
			num++
		}
	}
	return num, nil
}
