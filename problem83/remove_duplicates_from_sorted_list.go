package problem83

import (
	"github.com/lutepluto/leetcode/common"
)

func deleteDuplicates(head *common.ListNode) *common.ListNode {
	iterator := head
	for iterator != nil {
		if iterator.Next != nil && iterator.Next.Val == iterator.Val {
			iterator.Next = iterator.Next.Next
		} else {
			iterator = iterator.Next
		}
	}
	return head
}
