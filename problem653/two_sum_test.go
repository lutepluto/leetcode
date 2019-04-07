package problem653

import (
	"fmt"
	"math"
	"testing"

	"github.com/lutepluto/leetcode/common"
)

func TestFindTarget(t *testing.T) {
	// tree := common.BinaryTreeFromLevelTraversal(5, 3, 6, 2, 4, math.MinInt32, 7)
	// fmt.Printf("%v\n", findTarget(tree, 9))
	// fmt.Printf("%v\n", findTarget(tree, 28))
	fmt.Printf("%v\n", findTarget(common.BinaryTreeFromLevelTraversal(1, 0, 4, -2, math.MinInt32, 3), 7))
}
