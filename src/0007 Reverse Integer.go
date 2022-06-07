package src

import (
	"math"
)

func reverse(x int) int {
	sgn := 1
	if x < 0 {
		sgn = -1
		x = -x
	}

	res := 0
	for x > 0 {
		if res > math.MaxInt32/10 || res == math.MaxInt32/10 && x%10 > math.MaxInt32%10 {
			return 0
		}
		res = res*10 + x%10
		x /= 10
	}

	return sgn * res
}
