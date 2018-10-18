package problem20

func isValid(s string) bool {
	if s == "" {
		return true
	}

	stack := make([]string, 0)
	opened := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}
	closed := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	for _, c := range s {
		if _, ok := opened[string(c)]; ok {
			stack = append(stack, string(c))
		} else if v, ok := closed[string(c)]; ok {
			if stack[len(stack)-1] != v {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
