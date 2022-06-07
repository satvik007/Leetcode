package src

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	mp := make(map[rune]int)
	ans := 0

	last := -1
	for i, ch := range s {
		if val, ok := mp[ch]; ok && val > last {
			last = val
		}
		ans = max(ans, i-last)
		mp[ch] = i
	}

	return ans
}
