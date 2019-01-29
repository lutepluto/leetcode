package problem100

import (
	"fmt"
	"testing"

	"github.com/lutepluto/leetcode/common"
)

func TestIsSame(t *testing.T) {
	p := common.BinaryTreeFromLevelTraversal(1, 2, 3)
	q := common.BinaryTreeFromLevelTraversal(1, 2, 3)
	fmt.Printf("%v\n", isSameTree(p, q))
}
