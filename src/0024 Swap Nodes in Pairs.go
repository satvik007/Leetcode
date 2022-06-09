package src

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	t := head.Next

	head.Next = swapPairs(t.Next)
	t.Next = head
	return t
}
