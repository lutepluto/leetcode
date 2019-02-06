package problem108

import (
	"fmt"
	"testing"

	"github.com/lutepluto/leetcode/common"
)

func TestSortedArrayToBST(t *testing.T) {
	nums := []int{-10, -3, 0, 5, 9}
	tree := sortedArrayToBST(nums)
	fmt.Printf("%v\n", common.PreOrderTraversal(tree))
	fmt.Printf("%v\n", common.LevelOrderTraversal(tree))
}
