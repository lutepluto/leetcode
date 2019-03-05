package problem169

// Moore voting algorithm
func majorityElement(nums []int) int {
	majority := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			count++
			majority = nums[i]
		} else if majority == nums[i] {
			count++
		} else {
			count--
		}
	}
	return majority
}
