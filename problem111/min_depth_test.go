package problem111

import (
	"fmt"
	"math"
	"testing"

	"github.com/lutepluto/leetcode/common"
)

func TestMinDepth(t *testing.T) {
	tree := common.BinaryTreeFromLevelTraversal(3, 9, 20, math.MinInt32, math.MinInt32, 15, 7)
	fmt.Printf("%v\n", minDepth(tree))

	tree2 := common.BinaryTreeFromLevelTraversal(1, 2)
	fmt.Printf("%v\n", minDepth(tree2))
}
