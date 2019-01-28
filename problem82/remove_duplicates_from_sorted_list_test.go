package problem82

import (
	"testing"

	"github.com/lutepluto/leetcode/common"
)

func TestDeleteDuplicates(t *testing.T) {
	p1 := common.ConstructLinkedList(1, 2, 3, 3, 4, 4, 5)
	common.PrintLinkedList(deleteDuplicates(p1))

	p2 := common.ConstructLinkedList(1, 1)
	common.PrintLinkedList(deleteDuplicates(p2))
}
