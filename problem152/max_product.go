package problem152

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	maxSofar := nums[0]
	minSofar := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxSofar, minSofar = minSofar, maxSofar
		}

		maxSofar *= nums[i]
		minSofar *= nums[i]
		if maxSofar < nums[i] {
			maxSofar = nums[i]
		}
		if minSofar > nums[i] {
			minSofar = nums[i]
		}

		if maxSofar > max {
			max = maxSofar
		}
	}
	return max
}
