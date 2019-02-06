package problem108

import (
	"github.com/lutepluto/leetcode/common"
)

func sortedArrayToBST(nums []int) *common.TreeNode {
	low := 0
	high := len(nums) - 1
	return insert(nums, low, high)
}

func insert(nums []int, low, high int) *common.TreeNode {
	if low > high {
		return nil
	}
	middle := (high + low) / 2
	root := &common.TreeNode{Val: nums[middle]}
	root.Left = insert(nums, low, middle-1)
	root.Right = insert(nums, middle+1, high)
	return root
}
