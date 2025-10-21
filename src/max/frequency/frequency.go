package frequency

import "sort"

func maxFrequency(nums []int, k int, numOperations int) int {
	sort.Ints(nums)
	n := len(nums)
	var out int
	var l, r, i int
	for num := nums[0]; num <= nums[n-1]; num++ {
		min := num - k
		max := num + k
		var count int
		for i < n && nums[i] < num {
			i++
		}
		for i < n && nums[i] == num {
			i++
			count++
		}
		for l < n && nums[l] < min {
			l++
		}
		for r < n && nums[r] <= max {
			r++
		}
		temp := r - l
		if temp-count > numOperations {
			temp = count + numOperations
		}
		if temp > out {
			out = temp
		}
	}
	return out
}
