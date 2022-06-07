package src

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	var root, current *ListNode

	for l1 != nil || l2 != nil {
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		v := v1 + v2 + carry

		if root == nil {
			root = &ListNode{}
			current = root
		} else {
			current.Next = &ListNode{}
			current = current.Next
		}
		current.Val = v % 10
		carry = v / 10

	}

	if carry != 0 {
		current.Next = &ListNode{}
		current = current.Next
		current.Val = carry
	}

	return root
}
