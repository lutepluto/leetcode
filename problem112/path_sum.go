package problem112

import (
	"github.com/lutepluto/leetcode/common"
)

func hasPathSum(root *common.TreeNode, sum int) bool {
	return pathSum(root, 0, sum)
}

func pathSum(root *common.TreeNode, sum, target int) bool {
	if root == nil {
		return false
	}
	sum += root.Val
	if root.Left == nil && root.Right == nil {
		if sum > target {
			return false
		}
		if sum == target {
			return true
		}
	}
	leftHasPathSum := pathSum(root.Left, sum, target)
	rightHasPathSum := pathSum(root.Right, sum, target)
	return leftHasPathSum || rightHasPathSum
}
