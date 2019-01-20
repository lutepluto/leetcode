package problem26

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	length := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] > nums[i-1] {
					nums[i], nums[j] = nums[j], nums[i]
					length++
					break
				}
				if j == len(nums)-1 {
					return length
				}
			}
		} else {
			length++
		}
	}
	return length
}

func solution(nums []int) int {
	length := len(nums)
	if length < 2 {
		return length
	}
	i := 0
	for j := 1; j < length; j++ {
		if nums[i] != nums[j] {
			i = i + 1
			nums[i] = nums[j]
		}
	}
	return i + 1
}
