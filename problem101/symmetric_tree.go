package problem101

import (
	"github.com/lutepluto/leetcode/common"
)

func isSymmetric(root *common.TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricRecursively(root.Left, root.Right)
}

func isSymmetricRecursively(left, right *common.TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left != nil && right != nil {
		return left.Val == right.Val &&
			isSymmetricRecursively(left.Left, right.Right) &&
			isSymmetricRecursively(left.Right, right.Left)
	}
	return false
}
