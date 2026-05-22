func isPalindrome(s string) bool {

	l, r := 0, len(s)-1
	s = strings.ToLower(s)

	for l <= r {
		if !checkCharacter(s[l]) {
			l++
			continue
		}

		if !checkCharacter(s[r]) {
			r--
			continue
		}

		if s[l] != s[r] {
			return false
		} else {
			l++
			r--
		}
	}

	return true
}

func checkCharacter(ch byte) bool {

	if ch >= 48 && ch <= 57 || ch >= 65 && ch <= 90 || ch >= 97 && ch <= 122 {
		return true
	}
	return false
}