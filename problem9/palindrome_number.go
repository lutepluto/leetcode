package problem9

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	// reverse x
	tmp := x
	reversed := 0
	for tmp != 0 {
		reversed = reversed*10 + tmp%10
		tmp = tmp / 10
	}
	return reversed == x
}

func solution(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	rev := 0
	for x > rev {
		rev = rev*10 + x%10
		x = x / 10
	}
	return x == rev || x == rev/10
}
