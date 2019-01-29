package problem101

import (
	"fmt"
	"math"
	"testing"

	"github.com/lutepluto/leetcode/common"
)

func TestIsSymmetric(t *testing.T) {
	tree1 := common.BinaryTreeFromLevelTraversal(1, 2, 2, 3, 4, 4, 3)
	fmt.Printf("is symmetric: %v\n", isSymmetric(tree1))

	tree2 := common.BinaryTreeFromLevelTraversal(1, 2, 2, math.MinInt32, 3, math.MinInt32, 3)
	fmt.Printf("is symmetric: %v\n", isSymmetric(tree2))
}
