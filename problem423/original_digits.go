package problem423

import "strconv"

func originalDigits(s string) string {
	count := make([]int, 10)
	for _, c := range s {
		l := string(c)
		switch l {
		case "z":
			count[0]++ // 0
		case "w":
			count[2]++ // 2
		case "u":
			count[4]++ // 4
		case "x":
			count[6]++ // 6
		case "g":
			count[8]++ // 8
		case "h":
			count[3]++ // 8+3
		case "f":
			count[5]++ // 4+5
		case "s":
			count[7]++ // 7+6
		case "i":
			count[9]++ // 5+6+8+9
		case "o":
			count[1]++ // 0+1+2+4
		}
	}

	count[3] -= count[8]
	count[5] -= count[4]
	count[7] -= count[6]
	count[9] = count[9] - count[5] - count[6] - count[8]
	count[1] = count[1] - count[0] - count[2] - count[4]

	var digits string

	for i := 0; i < len(count); i++ {
		for j := 0; j < count[i]; j++ {
			digits += strconv.Itoa(i)
		}
	}

	return digits
}
