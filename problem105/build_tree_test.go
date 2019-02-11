package problem105

import (
	"fmt"
	"testing"

	"github.com/lutepluto/leetcode/common"
)

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)

	fmt.Printf("%#v\n", common.LevelOrderTraversal(root))
}
