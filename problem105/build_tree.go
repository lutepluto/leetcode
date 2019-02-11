package problem105

import (
	"github.com/lutepluto/leetcode/common"
)

func buildTree(preorder, inorder []int) *common.TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &common.TreeNode{Val: preorder[0]}
	var leftPreorder, rightPreorder, leftInorder, rightInorder []int
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			leftPreorder = preorder[1 : 1+i]
			leftInorder = inorder[:i]
			rightPreorder = preorder[1+i:]
			rightInorder = inorder[i+1:]
			break
		}
	}

	root.Left = buildTree(leftPreorder, leftInorder)
	root.Right = buildTree(rightPreorder, rightInorder)
	return root
}
