package src

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func longestPalindrome(s string) string {
	n := len(s)
	d := make([]int, n)
	l, r := 0, 0
	m1, m2 := 0, 0
	mv := 0

	for i := 0; i < n; i++ {
		if r > i {
			d[i] = max(0, min(r-i, d[l+r-i]))
		}

		for i+d[i] < n && i-d[i] >= 0 && s[i+d[i]] == s[i-d[i]] {
			d[i]++
		}

		if i+d[i] > r {
			l, r = i-d[i], i+d[i]
		}

		if 2*d[i]-1 > mv {
			m1, m2 = i-d[i]+1, i+d[i]
			mv = 2*d[i] - 1
		}
	}

	l, r = 0, 0
	for i := 0; i < n; i++ {
		if r > i {
			d[i] = max(0, min(r-i, d[l+r-i]))
		} else {
			d[i] = 0
		}

		for i+1+d[i] < n && i-d[i] >= 0 && s[i+d[i]+1] == s[i-d[i]] {
			d[i]++
		}

		if i+d[i] > r {
			l, r = i-d[i], i+d[i]
		}

		if 2*d[i] > mv {
			m1, m2 = i-d[i]+1, i+d[i]+1
			mv = 2 * d[i]
		}
	}

	return s[m1:m2]
}
