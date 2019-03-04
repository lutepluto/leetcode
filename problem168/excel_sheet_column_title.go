package problem168

func convertToTitle(n int) string {
	var chars = []string{
		"Z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y",
	}

	var title string
	for n > 0 {
		r := n % 26
		title = chars[r] + title
		n = n / 26
		if r == 0 {
			n--
		}
	}
	return title
}
