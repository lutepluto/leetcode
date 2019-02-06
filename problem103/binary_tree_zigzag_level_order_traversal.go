package problem103

import (
	"github.com/lutepluto/leetcode/common"
)

func zigzagLevelOrder(root *common.TreeNode) [][]int {
	zigzag := make([][]int, 0)
	if root == nil {
		return zigzag
	}

	list := []*common.TreeNode{root}
	direction := true

	for len(list) > 0 {
		len := len(list)
		levelOrder := make([]int, 0)

		for i := len - 1; i >= 0; i-- {
			node := list[i]
			levelOrder = append(levelOrder, node.Val)

			if direction {
				if node.Left != nil {
					list = append(list, node.Left)
				}
				if node.Right != nil {
					list = append(list, node.Right)
				}
			} else {
				if node.Right != nil {
					list = append(list, node.Right)
				}
				if node.Left != nil {
					list = append(list, node.Left)
				}
			}
		}

		direction = !direction
		list = list[len:]
		zigzag = append(zigzag, levelOrder)
	}

	return zigzag
}
