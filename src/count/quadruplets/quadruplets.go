package quadruplets

import "sort"

func countQuadruplets(nums []int) int64 {
	n := len(nums)
	suffix := make([]int64, n)

	for i := n - 2; i >= 0; i-- {
		a := nums[i]
		for j := i + 1; j < n; j++ {
			if nums[j] > a {
				suffix[i]++
			}
		}
	}
	var ret int64

	prefix := make([]int, 0, n)
	prefix = append(prefix, nums[0])
	for i := 1; i+2 < n; i++ {
		a := nums[i]
		var temp int
		for j := i + 1; j+1 < n; j++ {
			if a > nums[j] {
				index := sort.SearchInts(prefix, nums[j])
				if index > 0 {
					ret += int64(index) * (suffix[i] - int64(temp))
				}
			} else {
				temp++
			}
		}
		index := sort.SearchInts(prefix, a)
		if index < len(prefix) {
			prefix = append(prefix, 0)
			copy(prefix[index+1:], prefix[index:])
			prefix[index] = a
		} else {
			prefix = append(prefix, a)
		}
	}
	return ret
}
