package utils

import . "gcode/model"

func ArrayToListNode(arr []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, val := range arr {
		node := &ListNode{
			Val: val,
		}
		cur.Next = node
		cur = cur.Next
	}
	return dummy.Next
}

func ListNodeToArray(node *ListNode) []int {
	var res []int
	for node != nil {
		res = append(res, node.Val)
		node = node.Next
	}
	return res
}
