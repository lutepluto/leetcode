package problem21

import (
	"fmt"
	"testing"

	"github.com/lutepluto/leetcode/common"
)

func TestMergeTwoLists(t *testing.T) {
	var p1, p2 *common.ListNode

	p1 = common.ConstructLinkedList(1, 2, 4)
	p2 = common.ConstructLinkedList(1, 3, 4)
	fmt.Printf("Merged list: ")
	common.PrintLinkedList(mergeTwoLists(p1, p2))

	p1 = common.ConstructLinkedList(1, 4, 8)
	p2 = common.ConstructLinkedList(5, 6, 9)
	fmt.Printf("Recursively merged list: ")
	common.PrintLinkedList(mergeTwoListsRecursively(p1, p2))
}
