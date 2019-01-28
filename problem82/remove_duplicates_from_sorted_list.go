package problem82

import (
	"github.com/lutepluto/leetcode/common"
)

func deleteDuplicates(head *common.ListNode) *common.ListNode {
	if head == nil {
		return head
	}
	dummy := &common.ListNode{
		Next: head,
	}
	p := dummy
	c := head
	n := c.Next
	for c != nil && n != nil {
		if n.Val != c.Val {
			p = c
			c = n
		} else {
			for n != nil && n.Val == c.Val {
				n = n.Next
			}
			p.Next = n
			c = n
			if n == nil {
				break
			}
		}
		n = n.Next
	}
	return dummy.Next
}
