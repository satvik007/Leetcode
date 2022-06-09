package src

func romanToInt(s string) int {
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	mapping := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	ans := 0
	for i := 0; i < len(s); {
		for j := 0; j < len(roman); j++ {
			if i+len(roman[j]) <= len(s) && s[i:i+len(roman[j])] == roman[j] {
				ans += mapping[j]
				i += len(roman[j])
				break
			}
		}
	}

	return ans
}
