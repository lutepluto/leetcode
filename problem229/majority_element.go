package problem229

func majorityElements(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	var marjority1, marjority2 int
	count1, count2 := 0, 0

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if marjority1 == num {
			count1++
		} else if marjority2 == num {
			count2++
		} else if count1 == 0 {
			marjority1 = num
			count1++
		} else if count2 == 0 {
			marjority2 = num
			count2++
		} else {
			count1--
			count2--
		}
	}

	count1, count2 = 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == marjority1 {
			count1++
		} else if nums[i] == marjority2 {
			count2++
		}
	}

	result := make([]int, 0)
	if count1 > len(nums)/3 {
		result = append(result, marjority1)
	}
	if count2 > len(nums)/3 {
		result = append(result, marjority2)
	}

	return result
}
