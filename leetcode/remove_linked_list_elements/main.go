package remove_linked_list_elements

import . "gcode/model"

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	for tmp := dummy; tmp.Next != nil; {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return dummy.Next
}

func removeElements2(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}
