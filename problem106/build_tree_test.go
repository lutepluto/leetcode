package problem106

import (
	"fmt"
	"testing"

	"github.com/lutepluto/leetcode/common"
)

func TestBuildTree(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	root := buildTree(inorder, postorder)

	fmt.Printf("%#v\n", common.LevelOrderTraversal(root))
}
