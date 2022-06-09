package src

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var res, cur *ListNode

	for list1 != nil || list2 != nil {
		t := &ListNode{}
		if res == nil {
			res = t
		} else {
			cur.Next = t
		}
		cur = t

		if list1 == nil {
			cur.Val = list2.Val
			list2 = list2.Next
		} else if list2 == nil {
			cur.Val = list1.Val
			list1 = list1.Next
		} else if list1.Val < list2.Val {
			cur.Val = list1.Val
			list1 = list1.Next
		} else {
			cur.Val = list2.Val
			list2 = list2.Next
		}
	}

	return res
}
