package problem102

import (
	"fmt"
	"math"
	"testing"

	"github.com/lutepluto/leetcode/common"
)

func TestLevelOrder(t *testing.T) {
	tree := common.BinaryTreeFromLevelTraversal(3, 9, 20, math.MinInt32, math.MinInt32, 15, 7)
	levelOrders := levelOrder(tree)
	for _, levelOrder := range levelOrders {
		fmt.Printf("%#v\n", levelOrder)
	}
}
