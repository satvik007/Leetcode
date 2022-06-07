package src

import (
	"math"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	y := x
	res := 0
	for y > 0 {
		if res > math.MaxInt32/10 || res == math.MaxInt32/10 && y%10 > math.MaxInt32%10 {
			return false
		}
		res = res*10 + y%10
		y /= 10
	}

	return res == x
}
