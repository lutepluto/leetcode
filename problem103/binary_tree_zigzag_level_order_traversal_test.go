package problem103

import (
	"fmt"
	"math"
	"testing"

	"github.com/lutepluto/leetcode/common"
)

func TestZigzagLevelOrder(t *testing.T) {
	tree := common.BinaryTreeFromLevelTraversal(1, 2, 3, 4, math.MinInt32, math.MinInt32, 5)
	zigzag := zigzagLevelOrder(tree)

	for _, levelOrder := range zigzag {
		fmt.Printf("%#v\n", levelOrder)
	}
}
