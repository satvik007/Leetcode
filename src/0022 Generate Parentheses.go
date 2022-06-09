package src

func generateParenthesis(n int) []string {
	ans := []string{}

	var generate func(string, int)
	generate = func(s string, left int) {
		if len(s) == 2*n {
			if left == 0 {
				ans = append(ans, s)
			}
			return
		}

		generate(s+"(", left+1)
		if left > 0 {
			generate(s+")", left-1)
		}
	}

	generate("", 0)

	return ans
}
