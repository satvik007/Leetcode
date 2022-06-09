package src

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	ans := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			left, right := j+1, len(nums)-1

			for left <= right {
				mid := left + (right-left)/2
				if math.Abs(float64(nums[i]+nums[j]+nums[mid]-target)) < math.Abs(float64(ans-target)) {
					ans = nums[i] + nums[j] + nums[mid]
				}
				if nums[mid] > (target - nums[i] - nums[j]) {
					right = mid - 1
				} else if nums[mid] < (target - nums[i] - nums[j]) {
					left = mid + 1
				} else {
					break
				}
			}

		}
	}

	return ans
}
