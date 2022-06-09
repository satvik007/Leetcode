package src

import (
	"math"
)

func mergeKLists(lists []*ListNode) *ListNode {
	var res, cur *ListNode

	flag := true
	for flag {
		flag = false

		mini := -1
		mv := math.MaxInt32
		for i, list := range lists {
			if list == nil {
				continue
			}
			if list.Val < mv {
				mv = list.Val
				mini = i
				flag = true
			}
		}

		if !flag {
			break
		}

		lists[mini] = lists[mini].Next
		t := &ListNode{Val: mv}
		if res == nil {
			res = t
		} else {
			cur.Next = t
		}
		cur = t
	}

	return res
}
