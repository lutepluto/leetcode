package problem106

import (
	"github.com/lutepluto/leetcode/common"
)

func buildTree(inorder, postorder []int) *common.TreeNode {
	postLen := len(postorder)
	inLen := len(inorder)
	if inLen == 0 || postLen == 0 {
		return nil
	}

	root := &common.TreeNode{Val: postorder[postLen-1]}
	var leftInorder, rightInorder, leftPostorder, rightPostorder []int
	for i := 0; i < inLen; i++ {
		if inorder[i] == postorder[postLen-1] {
			leftInorder = inorder[:i]
			rightInorder = inorder[i+1:]
			leftPostorder = postorder[:i]
			rightPostorder = postorder[i : postLen-1]
		}
	}

	root.Left = buildTree(leftInorder, leftPostorder)
	root.Right = buildTree(rightInorder, rightPostorder)
	return root
}
