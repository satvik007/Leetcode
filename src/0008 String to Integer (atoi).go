package src

import (
	"math"
)

func myAtoi(s string) int {
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}

	sign := 1
	if i < len(s) && (s[i] == '-' || s[i] == '+') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	var res int

	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		if sign == -1 {
			if res > math.MaxInt32/10 || res == math.MaxInt32/10 && s[i]-'0' > (math.MaxInt32%10) {
				return math.MinInt32
			}
			res = res*10 + int(s[i]-'0')
		} else {
			if res > math.MaxInt32/10 || res == math.MaxInt32/10 && s[i]-'0' > (math.MaxInt32%10) {
				return math.MaxInt32
			}
			res = res*10 + int(s[i]-'0')
		}
		i++
	}

	return sign * res
}
