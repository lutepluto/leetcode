package problem67

var numbers = map[string]int{
	"0": 0,
	"1": 1,
}

func addBinary(a, b string) string {
	aLen := len(a)
	bLen := len(b)
	var total string
	extra := 0

	for i := 0; aLen-i > 0 || bLen-i > 0; i++ {
		aDigit := 0
		bDigit := 0
		if aLen-i > 0 {
			aDigit = numbers[string(a[aLen-i-1])]
		}
		if bLen-i > 0 {
			bDigit = numbers[string(b[bLen-i-1])]
		}
		sum := aDigit + bDigit + extra
		if sum%2 == 1 {
			total = "1" + total
		} else {
			total = "0" + total
		}
		extra = sum / 2
	}
	if extra != 0 {
		total = "1" + total
	}
	return total
}
