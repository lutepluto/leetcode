package problem100

import "github.com/lutepluto/leetcode/common"

func isSameTree(p, q *common.TreeNode) bool {
	if p != nil && q != nil {
		if p.Val == q.Val {
			isLeftSame := isSameTree(p.Left, q.Left)
			if isLeftSame {
				return isSameTree(p.Right, q.Right)
			}
			return isLeftSame
		}
		return false
	} else if p == nil && q == nil {
		return true
	}
	return false
}
