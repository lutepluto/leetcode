package problem112

import (
	"fmt"
	"math"
	"testing"

	"github.com/lutepluto/leetcode/common"
)

func TestHasPathSum(t *testing.T) {
	tree := common.BinaryTreeFromLevelTraversal(5, 4, 8, 11, math.MinInt32, 13, 4, 7, 2, math.MinInt32, math.MinInt32, math.MinInt32, 1)
	fmt.Printf("%#v\n", hasPathSum(tree, 22))

	tree2 := common.BinaryTreeFromLevelTraversal(1, 2)
	fmt.Printf("%#v\n", hasPathSum(tree2, 1))
}
