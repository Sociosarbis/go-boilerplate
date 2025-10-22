package frequency2

import "sort"

func maxFrequency(nums []int, k int, numOperations int) int {
	sort.Ints(nums)
	n := len(nums)
	var out int
	var l, r, i int
	options := make([]int, 0, 3*n)
	for i, num := range nums {
		if (i != 0 && num != nums[i-1]) || i == 0 {
			if k == 0 {
				options = append(options, num)
			} else {
				options = append(options, num-k, num, num+k)
			}
		}
	}
	sort.Ints(options)
	for _, num := range options {
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
