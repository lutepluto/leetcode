package problem104

import (
	"fmt"
	"math"
	"testing"

	"github.com/lutepluto/leetcode/common"
)

func TestMaxDepth(t *testing.T) {
	tree := common.BinaryTreeFromLevelTraversal(3, 9, 20, math.MinInt32, math.MinInt32, 15, 7)
	fmt.Printf("max depth: %d\n", maxDepth(tree))
}
