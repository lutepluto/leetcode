package problem653

import (
	"fmt"

	"github.com/lutepluto/leetcode/common"
)

func findTarget(root *common.TreeNode, k int) bool {
	numbers := bst2SortedArray(root, make([]int, 0))
	if len(numbers) == 0 {
		return false
	}
	fmt.Printf("%#v\n", numbers)
	for i := len(numbers) - 1; i > 0; i-- {
		if numbers[i] <= k-numbers[0] {
			j := 0
			for j < i {
				if numbers[j]+numbers[i] == k {
					return true
				} else if numbers[j]+numbers[i] < k {
					j++
				} else {
					i--
				}
			}
		}
	}
	return false
}

func bst2SortedArray(root *common.TreeNode, array []int) []int {
	if root == nil {
		return array
	}
	array = bst2SortedArray(root.Left, array)
	array = append(array, root.Val)
	return bst2SortedArray(root.Right, array)
}
