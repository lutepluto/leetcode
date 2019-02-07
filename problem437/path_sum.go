package problem437

import (
	"github.com/lutepluto/leetcode/common"
)

func pathSum(root *common.TreeNode, sum int) int {
	result := 0
	if root != nil {
		result += pathSum(root.Left, sum)
		result += pathSum(root.Right, sum)
		result += dfs(root, sum)
	}
	return result
}

func dfs(root *common.TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	result := 0
	if root.Val == sum {
		result = 1
	}
	return result + dfs(root.Left, sum-root.Val) + dfs(root.Right, sum-root.Val)
}
