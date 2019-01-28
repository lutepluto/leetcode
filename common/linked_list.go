package common

import "fmt"

// ListNode is the node of linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// ConstructLinkedList constructs linked list with given integers
func ConstructLinkedList(numbers ...int) *ListNode {
	head := &ListNode{}
	tmp := head
	for _, n := range numbers {
		tmp.Next = &ListNode{Val: n}
		tmp = tmp.Next
	}
	return head.Next
}

// PrintLinkedList prints linked list item one by one
func PrintLinkedList(l *ListNode) {
	for p := l; p != nil; p = p.Next {
		fmt.Printf("%d ", p.Val)
	}
	fmt.Printf("\n")
}
