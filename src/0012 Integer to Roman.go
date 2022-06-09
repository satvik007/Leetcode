package src

import (
	"strings"
)

func intToRoman(num int) string {
	roman := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	mapping := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}

	ans := ""

	for i := len(roman) - 1; i >= 0; i-- {
		if roman[i] <= num {
			v := num / roman[i]
			ans += strings.Repeat(mapping[i], v)
			num -= v * roman[i]
		}
	}

	return ans
}
