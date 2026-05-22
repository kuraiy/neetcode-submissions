func isValid(s string) bool {
	openedBrackets := make([]rune, 0)

	for _, v := range s {
		if v == '[' || v == '(' || v == '{' {
			openedBrackets = append(openedBrackets, v)
		} else {
			n := len(openedBrackets) - 1
			if n == -1 {
				return false
			}
			switch v {
			case ']':
				if openedBrackets[n] == '[' {
					openedBrackets = openedBrackets[:n]
				} else {
					return false
				}
			case '}':
				if openedBrackets[n] == '{' {
					openedBrackets = openedBrackets[:n]
				} else {
					return false
				}
			case ')':
				if openedBrackets[n] == '(' {
					openedBrackets = openedBrackets[:n]
				} else {
					return false
				}
			}
		}
	}

	if len(openedBrackets) == 0 {
		return true
	} else {
		return false
	}
}