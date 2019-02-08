package problem125

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		for !isAlphanumeric(rune(s[i])) && i < j {
			i++
		}
		for !isAlphanumeric(rune(s[j])) && j > i {
			j--
		}
		if i >= j {
			break
		}
		si := toUppercase(rune(s[i]))
		sj := toUppercase(rune(s[j]))
		if si != sj {
			return false
		}
		i++
		j--
	}
	return true
}

func toUppercase(c rune) rune {
	if c >= 'a' && c <= 'z' {
		return c - 'a' + 'A'
	}
	return c
}

func isAlphanumeric(c rune) bool {
	if c < '0' || (c > '9' && c < 'A') || (c > 'Z' && c < 'a') || c > 'z' {
		return false
	}
	return true
}
