package problem110

import (
	"github.com/lutepluto/leetcode/common"
)

func isBalanced(root *common.TreeNode) bool {
	if root == nil {
		return true
	}
	depth := findDepth(root)
	return depth != -1
}

func findDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := findDepth(root.Left)
	if leftDepth == -1 {
		return -1
	}
	rightDepth := findDepth(root.Right)
	if rightDepth == -1 {
		return -1
	}

	var diff, max int
	if leftDepth > rightDepth {
		diff = leftDepth - rightDepth
		max = leftDepth
	} else {
		diff = rightDepth - leftDepth
		max = rightDepth
	}

	if diff > 1 {
		return -1
	}
	return max + 1
}
