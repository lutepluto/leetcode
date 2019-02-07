package problem437

import (
	"fmt"
	"math"
	"testing"

	"github.com/lutepluto/leetcode/common"
)

func TestPathSum(t *testing.T) {
	tree := common.BinaryTreeFromLevelTraversal(
		// 10, 5, -3, 3, 2, math.MinInt32, 11, 3, -2, math.MinInt32, 1,
		// 5, 3, 2, 3, -2, math.MinInt32, 1,
		// -3, math.MinInt32, 11,
		1, -2, -3, 1, 3, -2, math.MinInt32, -1,
	)
	fmt.Printf("%#v\n", pathSum(tree, -1))
}
