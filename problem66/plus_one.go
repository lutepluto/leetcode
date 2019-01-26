package problem66

func plusOne(digits []int) []int {
	operand := 1
	result := make([]int, len(digits))
	for i := len(digits) - 1; i >= 0; i-- {
		total := digits[i] + operand
		result[i] = total % 10
		operand = total / 10
	}
	if operand != 0 {
		result = append([]int{operand}, result...)
	}
	return result
}
