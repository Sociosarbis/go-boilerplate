package ii

import "sort"

func abs(num int) int {
	if num > 0 {
		return num
	}
	return -num
}

func getDiff(l1, r1, l2, r2 int) int {
	var min int
	if l2 < l1 {
		min = l2
	} else {
		min = l1
	}
	if r2 > r1 {
		return r2 - min
	} else {
		return r1 - min
	}
}

func smallestRangeII(nums []int, k int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	sort.Ints(nums)
	ret := nums[n-1] - nums[0]
	if k == 0 {
		return ret
	}
	if n == 2 {
		return abs(ret - 2*k)
	}
	end := n - 1
	l := nums[0] + k
	r := nums[n-1] - k
	for i := 0; i < end; i++ {
		temp := getDiff(l, nums[i]+k, nums[i+1]-k, r)
		if temp < ret {
			ret = temp
		}
	}
	return ret
}
