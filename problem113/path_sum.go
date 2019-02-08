package problem113

import (
	"github.com/lutepluto/leetcode/common"
)

func pathSum(root *common.TreeNode, sum int) [][]int {
	return doPathSum(root, sum, make([]int, 0), make([][]int, 0))
}

func doPathSum(root *common.TreeNode, sum int, path []int, paths [][]int) [][]int {
	if root == nil {
		return paths
	}
	if root.Left == nil && root.Right == nil && sum == root.Val {
		pathCopy := make([]int, len(path)+1)
		for idx, value := range path {
			pathCopy[idx] = value
		}
		pathCopy[len(path)] = root.Val
		paths = append(paths, pathCopy)
	} else {
		paths = doPathSum(root.Left, sum-root.Val, append(path, root.Val), paths)
		paths = doPathSum(root.Right, sum-root.Val, append(path, root.Val), paths)
	}
	return paths
}
