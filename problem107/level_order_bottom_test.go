package problem107

import (
	"fmt"
	"math"
	"testing"

	"github.com/lutepluto/leetcode/common"
)

func TestLevelOrderBottom(t *testing.T) {
	tree := common.BinaryTreeFromLevelTraversal(3, 9, 20, math.MinInt32, math.MinInt32, 15, 7)
	levelOrder := levelOrderBottom(tree)
	for _, level := range levelOrder {
		fmt.Printf("%#v\n", level)
	}
}
