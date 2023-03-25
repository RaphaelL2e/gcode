package linked_list_cycle

import (
	. "gcode/model"
)

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */
func hasCycle(head *ListNode) bool {
	set := make(map[*ListNode]bool)
	for head != nil {
		if _, ok := set[head]; ok {
			return true
		} else {
			set[head] = true
		}
		head = head.Next
	}
	return false
}

func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next
	}
	return true
}
