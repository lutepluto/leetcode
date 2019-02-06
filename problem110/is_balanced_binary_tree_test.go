package problem110

import (
	"fmt"
	"math"
	"testing"

	"github.com/lutepluto/leetcode/common"
)

func TestIsBalanced(t *testing.T) {
	tree1 := common.BinaryTreeFromLevelTraversal(3, 9, 20, math.MinInt32, math.MinInt32, 15, 7)
	fmt.Printf("%v\n", isBalanced(tree1))

	tree2 := common.BinaryTreeFromLevelTraversal(1, 2, 2, 3, 3, math.MinInt32, math.MinInt32, 4, 4)
	fmt.Printf("%v\n", isBalanced(tree2))

	tree3 := common.BinaryTreeFromLevelTraversal(1, math.MinInt32, 2, math.MinInt32, math.MinInt32, math.MinInt32, 3)
	fmt.Printf("%v\n", isBalanced(tree3))

}
