package problem107

import (
	"github.com/lutepluto/leetcode/common"
)

func levelOrderBottom(root *common.TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*common.TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		level := make([]int, 0)
		newStack := make([]*common.TreeNode, 0)
		for _, node := range stack {
			level = append(level, node.Val)
			if node.Left != nil {
				newStack = append(newStack, node.Left)
			}
			if node.Right != nil {
				newStack = append(newStack, node.Right)
			}
		}

		result = append(result, level)
		stack = newStack
	}

	reversed := make([][]int, len(result))
	for i := len(result) - 1; i >= 0; i-- {
		reversed[i] = result[len(result)-1-i]
	}
	return reversed
}
