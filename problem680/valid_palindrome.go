package problem680

func validPalindrome(s string) bool {
	idx := isPalindrome(s)
	if idx >= 0 && isPalindrome(s[idx+1:len(s)-idx]) >= 0 &&
		isPalindrome(s[idx:len(s)-1-idx]) >= 0 {
		return false
	}
	return true
}

func isPalindrome(s string) int {
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return i
		}
		i++
		j--
	}
	return -1
}
