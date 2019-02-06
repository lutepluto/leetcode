package problem102

import (
	"github.com/lutepluto/leetcode/common"
)

func levelOrder(root *common.TreeNode) [][]int {
	levelOrders := make([][]int, 0)
	if root == nil {
		return levelOrders
	}

	queue := []*common.TreeNode{root}
	for len(queue) > 0 {
		oldLength := len(queue)
		levelOrder := make([]int, 0)
		for i := 0; i < oldLength; i++ {
			node := queue[i]
			levelOrder = append(levelOrder, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[oldLength:]
		levelOrders = append(levelOrders, levelOrder)
	}
	return levelOrders
}
