package src

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	ans := make([][]int, 0)

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j-1] == nums[j] {
				continue
			}
			for k := j + 1; k < len(nums)-1; k++ {
				if k > j+1 && nums[k-1] == nums[k] {
					continue
				}
				left, right := k+1, len(nums)-1
				for left <= right {
					mid := left + (right-left)/2
					v := target - nums[i] - nums[j] - nums[k]
					if nums[mid] > v {
						right = mid - 1
					} else if nums[mid] < v {
						left = mid + 1
					} else {
						ans = append(ans, []int{nums[i], nums[j], nums[k], nums[mid]})
						break
					}
				}
			}
		}
	}

	return ans
}
