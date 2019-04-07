package problem167

func twoSum(numbers []int, target int) []int {
	for i := len(numbers) - 1; i >= 0; i-- {
		if numbers[i] <= target-numbers[0] {
			j := 0
			for j < i {
				if numbers[j]+numbers[i] == target {
					return []int{j + 1, i + 1}
				} else if numbers[j]+numbers[i] < target {
					j++
				} else {
					i--
				}
			}
		}
	}
	return nil
}
