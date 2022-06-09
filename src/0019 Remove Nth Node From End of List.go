package src

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	head, _ = traverse(head, nil, head, n)
	return head
}

func traverse(head *ListNode, prev *ListNode, current *ListNode, n int) (*ListNode, int) {
	if current == nil {
		return head, 1
	}

	head, count := traverse(head, current, current.Next, n)
	if count == n {
		if prev != nil {
			prev.Next = current.Next
		} else {
			head = current.Next
		}
	}

	return head, count + 1
}
