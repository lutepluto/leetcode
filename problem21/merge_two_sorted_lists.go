package problem21

import "github.com/lutepluto/leetcode/common"

func mergeTwoLists(l1, l2 *common.ListNode) *common.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head, tmp *common.ListNode
	if l1.Val > l2.Val {
		head = l2
		l2 = l2.Next
	} else {
		head = l1
		l1 = l1.Next
	}
	p1 := l1
	p2 := l2
	tmp = head
	for p1 != nil || p2 != nil {
		if p1 == nil {
			tmp.Next = p2
			break
		} else if p2 == nil {
			tmp.Next = p1
			break
		} else if p1.Val > p2.Val {
			tmp.Next = p2
			p2 = p2.Next
		} else {
			tmp.Next = p1
			p1 = p1.Next
		}
		tmp = tmp.Next
	}
	return head
}

func mergeTwoListsRecursively(l1, l2 *common.ListNode) *common.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l2.Next = mergeTwoListsRecursively(l1, l2.Next)
		return l2
	}
	l1.Next = mergeTwoListsRecursively(l1.Next, l2)
	return l1
}
