package src

func letterCombinations(digits string) []string {
	mp := []string{"", "", "abc", "def", "ghi", "jkl", "mno",
		"pqrs", "tuv", "wxyz"}

	old := []string{}
	new := []string{}

	for _, digit := range digits {
		for _, letter := range mp[digit-'0'] {
			if len(old) == 0 {
				new = append(new, string(letter))
			} else {
				for _, oldLetter := range old {
					new = append(new, oldLetter+string(letter))
				}
			}
		}
		old = new
		new = []string{}
	}

	return old
}
