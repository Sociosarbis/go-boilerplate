package elements

import "sort"

func maxDistinctElements(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	prev := nums[0] - k
	count := 1
	for i := 1; i < n; i++ {
		if nums[i] > prev {
			diff := nums[i] - prev - 1
			if diff > k {
				diff = k
			}
			prev = nums[i] - diff
			count++
		} else {
			diff := prev + 1 - nums[i]
			if diff <= k {
				prev = nums[i] + diff
				count++
			}
		}
	}
	return count
}
