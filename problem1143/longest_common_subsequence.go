package problem1143

func longestCommonSubsequence(text1 string, text2 string) int {
	len1 := len(text1)
	len2 := len(text2)
	if len1 == 0 || len2 == 0 {
		return 0
	}
	if text1[len1-1] == text2[len2-1] {
		return 1 + longestCommonSubsequence(text1[:len1-1], text2[:len2-1])
	}
	long1 := longestCommonSubsequence(text1[:len1-1], text2[:len2])
	long2 := longestCommonSubsequence(text1[:len1], text2[:len2-1])
	if long1 > long2 {
		return long1
	}
	return long2
}

func lcs(t1 string, t2 string) int {
	l1 := len(t1)
	l2 := len(t2)
	m := make([][]int, 0)
	for i := 0; i <= l1; i++ {
		r := make([]int, 0)
		for j := 0; j <= l2; j++ {
			if i == 0 || j == 0 {
				r = append(r, 0)
			} else if t1[i-1] == t2[j-1] {
				r = append(r, m[i-1][j-1]+1)
			} else {
				max1 := m[i-1][j]
				max2 := r[j-1]
				if max1 > max2 {
					r = append(r, max1)
				} else {
					r = append(r, max2)
				}
			}
		}
		m = append(m, r)
	}
	return m[l1][l2]
}
