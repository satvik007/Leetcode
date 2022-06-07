package src

import (
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	res := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = []byte(strings.Repeat(" ", len(s)))
	}

	x, y := 0, 0
	dir := 1
	for i := 0; i < len(s); i++ {
		res[x][y] = s[i]
		if x == numRows-1 {
			x--
			y++
			dir = -1
		} else if x == 0 {
			dir = 1
			x++
		} else {
			if dir == 1 {
				x++
			} else {
				x--
				y++
			}
		}
	}

	ans := ""
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(s); j++ {
			if res[i][j] != ' ' {
				ans += string(res[i][j])
			}
		}
	}

	return ans
}
