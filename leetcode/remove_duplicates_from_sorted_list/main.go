package remove_duplicates_from_sorted_list

import . "gcode/model"

func FailedDeleteDuplicates(head *ListNode) *ListNode {
	prev := &ListNode{Next: head}
	for tmp := prev; tmp != nil; {
		if tmp.Val == tmp.Next.Val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return prev.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	curr := head
	for curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return curr
}
