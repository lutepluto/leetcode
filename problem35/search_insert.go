package problem35

func searchInsert(nums []int, target int) int {
	return findIndex(nums, target, 0, len(nums)-1)
}

func findIndex(nums []int, target, low, high int) int {
	if low < high {
		mid := (low + high) / 2
		if nums[mid] > target {
			return findIndex(nums, target, low, mid-1)
		} else if nums[mid] == target {
			return mid
		} else {
			return findIndex(nums, target, mid+1, high)
		}
	}
	if nums[low] >= target {
		return low
	}
	return low + 1
}
