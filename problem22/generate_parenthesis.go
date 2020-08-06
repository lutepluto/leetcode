package problem22

func generateParenthesis(n int) []string {
	r := make([]string, 0)
	// return dfs(n*2, 0, 0, "", r)
	return dp("", 0, n*2, r)
}

// my solution
func dfs(n, idx, value int, str string, result []string) []string {
	if idx < n {
		for _, c := range "()" {
			v := 0

			if string(c) == "(" {
				v = 1
			} else if string(c) == ")" {
				v = -1
			}

			if value+v >= 0 {
				result = dfs(n, idx+1, value+v, str+string(c), result)
			}
		}
	} else if value == 0 {
		if str != "" {
			return append(result, str)
		}
	}
	return result
}

// better solution
func dp(prefix string, depth, n int, result []string) []string {
	if depth == 0 && n == 0 {
		return append(result, prefix)
	}
	if depth < n {
		result = dp(prefix+"(", depth+1, n-1, result)
	}
	if depth > 0 {
		result = dp(prefix+")", depth-1, n-1, result)
	}
	return result
}
