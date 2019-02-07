package problem113

import (
	"fmt"
	"math"
	"testing"

	"github.com/lutepluto/leetcode/common"
)

func TestPathSum(t *testing.T) {
	tree := common.BinaryTreeFromLevelTraversal(5, 4, 8, 11, math.MinInt32, 13, 4, 7, 2, math.MinInt32, math.MinInt32, math.MinInt32, math.MinInt32, 5, 1)
	// tree := common.BinaryTreeFromLevelTraversal(-2, math.MinInt32, -3)

	paths := pathSum(tree, 22)
	for _, path := range paths {
		fmt.Printf("%#v\n", path)
	}
}
