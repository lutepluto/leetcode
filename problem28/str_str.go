package problem28

func strStr(haystack, needle string) int {
	if needle == "" {
		return 0
	}
	haystackLen := len(haystack)
	needleLen := len(needle)
	if haystackLen < needleLen {
		return -1
	}
	i := 0
	j := 0
	for i < haystackLen && j < needleLen {
		if haystack[i] != needle[j] {
			i = i - j + 1
			j = 0
		} else {
			i++
			j++
		}
	}
	if i-j >= 0 && j == needleLen {
		return i - j
	}
	return -1
}
