package problem27

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	j := len(nums) - 1
	for {
		if nums[i] != val {
			i++
		}
		if nums[j] == val {
			j--
		}
		if i > j {
			break
		}
		if nums[i] == val && nums[j] != val {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return j + 1
}
