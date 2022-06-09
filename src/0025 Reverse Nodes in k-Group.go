package src

func reverseKGroup(head *ListNode, k int) *ListNode {
	cur := head
	val := make([]*ListNode, k)

	for i := 0; i < k; i++ {
		if cur == nil {
			return head
		}
		val[i] = cur
		cur = cur.Next
	}

	for i := k - 1; i > 0; i-- {
		val[i].Next = val[i-1]
	}
	val[0].Next = reverseKGroup(cur, k)

	return val[k-1]
}
