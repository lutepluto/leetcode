package problem53

func maxSubArray(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	max := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		n := nums[i]
		if sum < 0 && sum < n {
			sum = n
			if max < sum {
				max = sum
			}
		} else if n > 0 || sum > 0 {
			sum += n
			if max < sum {
				max = sum
			}
		}
	}
	return max
}

func maxSubArray2(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
	}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}
