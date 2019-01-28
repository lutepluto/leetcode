package problem83

import (
	"fmt"
	"testing"

	"github.com/lutepluto/leetcode/common"
)

func TestDeleteDuplicates(t *testing.T) {
	l1 := common.ConstructLinkedList(1, 1, 2)
	fmt.Printf("After removing duplicates: ")
	common.PrintLinkedList(deleteDuplicates(l1))

	l2 := common.ConstructLinkedList(1, 1, 2, 3, 3)
	fmt.Printf("After removing duplicates: ")
	common.PrintLinkedList(deleteDuplicates(l2))
}
