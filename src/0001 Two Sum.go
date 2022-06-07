package src

import (
	"sort"
)

func twoSum(nums []int, target int) []int {

	var a [][]int

	for i, v := range nums {
		a = append(a, []int{v, i})
	}

	sort.SliceStable(a, func(i int, j int) bool {
		return a[i][0] < a[j][0]
	})

	for i := 0; i < len(nums); i++ {
		left, right, mid := i+1, len(nums)-1, -1
		found := false

		for left <= right {
			mid = left + (right-left)/2

			if a[mid][0] > target-a[i][0] {
				right = mid - 1
			} else if a[mid][0] < target-a[i][0] {
				left = mid + 1
			} else {
				found = true
				break
			}
		}

		if found {
			return []int{a[i][1], a[mid][1]}
		}
	}

	return []int{-1, -1}
}
