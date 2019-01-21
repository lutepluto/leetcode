package problem38

import "strconv"

func countAndSay(n int) string {
	return doCount(n)
}

func doCount(n int) string {
	if n == 1 {
		return "1"
	}
	say := doCount(n - 1)
	i := 0
	var s []byte
	var result string
	for i < len(say) {
		if len(s) > 0 && s[len(s)-1] != say[i] {
			result += strconv.Itoa(len(s)) + string(s[len(s)-1])
			s = make([]byte, 0)
		} else {
			s = append(s, say[i])
			i++
		}
	}
	if len(s) > 0 {
		result += strconv.Itoa(len(s)) + string(s[len(s)-1])
	}
	return result
}
