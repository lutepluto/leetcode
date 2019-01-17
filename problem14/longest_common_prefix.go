package problem14

import "strings"

func horizontalScanning(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return prefix
			}
		}
	}
	return prefix
}

func verticalScanning(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}
	for i, b := range strs[0] {
		c := string(b)
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || c != string(strs[j][i]) {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func divideAndConquer(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}
	return longestCommonPrefix(strs, 0, len(strs)-1)
}

func longestCommonPrefix(strs []string, start, end int) string {
	if start == end {
		return strs[start]
	}
	mid := (start + end) / 2
	left := longestCommonPrefix(strs, start, mid)
	right := longestCommonPrefix(strs, mid+1, end)
	return mergeLongestCommonPrefix(left, right)
}

func mergeLongestCommonPrefix(left, right string) string {
	var minLen int
	if len(left) > len(right) {
		minLen = len(right)
	} else {
		minLen = len(left)
	}
	for i := 0; i < minLen; i++ {
		if left[i] != right[i] {
			return left[:i]
		}
	}
	return left[:minLen]
}
