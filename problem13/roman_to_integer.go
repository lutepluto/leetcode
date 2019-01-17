package problem13

var digits = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		c := string(s[i])
		current := digits[c]
		if i < len(s)-1 {
			d := string(s[i+1])
			if digits[d] > digits[c] {
				current = digits[d] - digits[c]
				i = i + 1
			}
		}
		result = result + current
	}
	return result
}
