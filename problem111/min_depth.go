package problem111

import (
	"github.com/lutepluto/leetcode/common"
)

func minDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}
	leftDepth := minDepth(root.Left)
	rightDepth := minDepth(root.Right)
	if leftDepth > rightDepth {
		return rightDepth + 1
	}
	return leftDepth + 1
}
