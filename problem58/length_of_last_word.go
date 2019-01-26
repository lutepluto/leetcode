package problem58

func lengthOfLastWord(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}
	count := 0
	for i := length - 1; i >= 0; i-- {
		c := string(s[i])
		if count != 0 && c == " " {
			break
		}
		if c != " " {
			count++
		}
	}
	return count
}
