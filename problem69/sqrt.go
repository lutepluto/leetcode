package problem69

func sqrt(x int) int {
	low := 0
	high := x
	for low <= high {
		mid := (low + high) / 2
		if mid*mid == x {
			return mid
		}
		if mid*mid > x {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low - 1
}
