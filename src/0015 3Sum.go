package src

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	ans := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for j := i + 1; j < len(nums)-1; j++ {
			if j > i+1 && nums[j-1] == nums[j] {
				continue
			}

			left, right, match := j+1, len(nums)-1, -1

			for left <= right {
				mid := left + (right-left)/2
				if nums[mid] > (-nums[i] - nums[j]) {
					right = mid - 1
				} else if nums[mid] < (-nums[i] - nums[j]) {
					left = mid + 1
				} else {
					match = mid
					break
				}
			}

			if match != -1 {
				v := []int{nums[i], nums[j], nums[match]}
				ans = append(ans, v)
			}
		}
	}

	return ans
}
