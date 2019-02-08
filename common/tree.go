package common

import (
	"math"
	"strconv"
)

// TreeNode is a node for binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BinaryTreeFromLevelTraversal constructs binary tree from level
// order traversal node values
func BinaryTreeFromLevelTraversal(values ...int) *TreeNode {
	return fromLevelTraversal(values, 0, len(values))
}

func fromLevelTraversal(values []int, idx, length int) *TreeNode {
	if idx >= length {
		return nil
	}
	if values[idx] == math.MinInt32 {
		return nil
	}
	node := &TreeNode{Val: values[idx]}
	node.Left = fromLevelTraversal(values, 2*idx+1, length)
	node.Right = fromLevelTraversal(values, 2*idx+2, length)
	return node
}

// LevelOrderTraversal traverse binary tree using level order
func LevelOrderTraversal(root *TreeNode) []string {
	values := make([]string, 0)
	if root == nil {
		return values
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left != nil || node.Right != nil {
			if node.Left == nil {
				queue = append(queue, &TreeNode{Val: math.MinInt32})
			} else {
				queue = append(queue, node.Left)
			}

			if node.Right == nil {
				queue = append(queue, &TreeNode{Val: math.MinInt32})
			} else {
				queue = append(queue, node.Right)
			}
		}

		if node.Val > math.MinInt32 {
			values = append(values, strconv.Itoa(node.Val))
		} else {
			values = append(values, "null")
		}
	}
	return values
}

// PreOrderTraversal traverse binary tree using pre-order
func PreOrderTraversal(root *TreeNode) []int {
	values := make([]int, 0)
	return preOrderTraversal(root, values)
}

func preOrderTraversal(root *TreeNode, values []int) []int {
	if root != nil {
		values = append(values, root.Val)
		values = preOrderTraversal(root.Left, values)
		values = preOrderTraversal(root.Right, values)
	}
	return values
}
