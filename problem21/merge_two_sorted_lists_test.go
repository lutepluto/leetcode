package problem21

import (
	"fmt"
	"testing"
)

func constructList(numbers ...int) *ListNode {
	head := &ListNode{}
	tmp := head
	for _, n := range numbers {
		tmp.Next = &ListNode{Val: n}
		tmp = tmp.Next
	}
	return head.Next
}

func printList(l *ListNode) {
	for p := l; p != nil; p = p.Next {
		fmt.Printf("%d ", p.Val)
	}
	fmt.Printf("\n")
}

func TestMergeTwoLists(t *testing.T) {
	var p1, p2 *ListNode

	p1 = constructList(1, 2, 4)
	p2 = constructList(1, 3, 4)
	fmt.Printf("Merged list: ")
	printList(mergeTwoLists(p1, p2))

	p1 = constructList(1, 4, 8)
	p2 = constructList(5, 6, 9)
	fmt.Printf("Recursively merged list: ")
	printList(mergeTwoListsRecursively(p1, p2))
}
