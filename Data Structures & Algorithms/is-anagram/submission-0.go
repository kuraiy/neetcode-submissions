func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	ms, mt := make(map[byte]int), make(map[byte]int)

	for i := range s {
		ms[s[i]]++
		mt[t[i]]++
	}

	for key, _ := range ms {
		if ms[key] != mt[key] {
			return false
		}
	}
	return true
}
