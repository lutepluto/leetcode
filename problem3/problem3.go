package problem3

func lengthOfLongestSubstring(s string) int {
	bucket := make(map[string]int)
	var start, end, max int
	for i, runeValue := range s {
		if i == 0 {
			start = i
		}
		end = i
		if idx, ok := bucket[string(runeValue)]; ok {
			if start <= idx {
				start = idx + 1
			}
		}
		bucket[string(runeValue)] = i
		if max < end-start+1 {
			max = end - start + 1
		}
	}
	return max
}
