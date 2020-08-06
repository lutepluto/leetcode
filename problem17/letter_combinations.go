package problem17

var m = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	r := make([]string, 0)
	dfs(digits, 0, "", &r)
	return r
}

func dfs(digits string, digitIndex int, str string, result *[]string) {
	if digitIndex < len(digits) {
		letters := m[string(digits[digitIndex])]
		for _, letter := range letters {
			dfs(digits, digitIndex+1, str+letter, result)
		}
	} else if str != "" {
		*result = append(*result, str)
	}
}
