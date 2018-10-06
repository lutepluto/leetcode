package atoi

import "math"

func atoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	sign := 1
	var result int32

	i := 0
	for ; i < len(str); i++ {
		if string(str[i]) != " " {
			break
		}
	}
	if i >= len(str) {
		return 0
	}

	char := string(str[i])
	if char == "+" || char == "-" {
		if char == "-" {
			sign = -1
		}
		i++
	}

	for ; i < len(str); i++ {
		char = string(str[i])
		if char >= "0" && char <= "9" {
			digit := rune(str[i]) - []rune("0")[0]
			if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > math.MaxInt32%10) {
				if sign == 1 {
					return math.MaxInt32
				} else if sign == -1 {
					return math.MinInt32
				}
			}
			result = result*10 + digit
		} else {
			break
		}
	}
	return sign * int(result)
}
