package problem141

import (
	"fmt"
	"testing"

	"github.com/lutepluto/leetcode/common"
)

func TestHasCycle(t *testing.T) {
	head := &common.ListNode{Val: 3}
	head.Next = &common.ListNode{Val: 2}
	head.Next.Next = &common.ListNode{Val: 0}
	head.Next.Next.Next = &common.ListNode{Val: -4}
	head.Next.Next.Next.Next = head.Next

	fmt.Printf("%v\n", hasCycle(head))
}
