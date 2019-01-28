package problem100

import (
	"fmt"
	"testing"

	"github.com/lutepluto/leetcode/common"
)

func TestIsSame(t *testing.T) {
	p := &common.TreeNode{Val: 1}
	p.Left = &common.TreeNode{Val: 2}
	p.Right = &common.TreeNode{Val: 3}

	q := &common.TreeNode{Val: 1}
	q.Left = &common.TreeNode{Val: 2}
	q.Right = &common.TreeNode{Val: 3}

	fmt.Printf("%v\n", isSameTree(q, q))
}
