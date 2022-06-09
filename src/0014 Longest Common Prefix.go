package src

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestCommonPrefix(strs []string) string {
	var left, right, ans int
	right = len(strs[0])
	for _, str := range strs {
		right = min(right, len(str))
	}

	for left <= right {
		mid := left + (right-left)/2

		equal := true
		for i := 1; i < len(strs); i++ {
			for j := 0; j < mid; j++ {
				if strs[0][j] != strs[i][j] {
					equal = false
					break
				}
			}
		}

		if equal {
			ans = max(ans, mid)
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return strs[0][:ans]
}
