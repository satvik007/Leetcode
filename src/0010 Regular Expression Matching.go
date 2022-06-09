package src

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	if len(p) >= 2 && p[1] == '*' {
		if isMatch(s, p[2:]) {
			return true
		}
		for i := 0; i < len(s) && (s[i] == p[0] || p[0] == '.'); i++ {
			if isMatch(s[i+1:], p[2:]) {
				return true
			}
		}
	}

	if len(s) == 0 {
		return false
	}

	if s[0] == p[0] || p[0] == '.' {
		return isMatch(s[1:], p[1:])
	}

	return false
}
