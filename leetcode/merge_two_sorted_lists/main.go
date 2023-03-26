package merge_two_sorted_lists

import . "gcode/model"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	prev := dummy
	for ; list1 != nil && list2 != nil; prev = prev.Next {
		if list1.Val <= list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else if list1.Val >= list2.Val {
			prev.Next = list2
			list2 = list2.Next
		}
	}
	if list1 != nil {
		prev.Next = list1
	}
	if list2 != nil {
		prev.Next = list2
	}
	return dummy.Next
}

func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val > list2.Val {
		return mergeTwoLists2(list2, list1)
	}
	p, q := list1, list2
	for p.Next != nil && q != nil {
		for p.Next != nil && p.Val <= q.Val && p.Next.Val <= q.Val {
			p = p.Next
		}
		if p.Next != nil {
			node := &ListNode{Val: q.Val}
			node.Next = p.Next
			p.Next = node

			q = q.Next
			p = p.Next
		}
	}
	if p.Next == nil {
		p.Next = q
	}
	return list1
}
