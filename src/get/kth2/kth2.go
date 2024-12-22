package kth2

import "sort"

func resolve(dp map[int]int, num int) int {
	if num == 1 {
		return 0
	}
	if v, ok := dp[num]; ok {
		return v
	}
	if num%2 == 0 {
		dp[num] = resolve(dp, num/2) + 1
	} else {
		dp[num] = resolve(dp, num*3+1) + 1
	}
	return dp[num]
}

func getKth(lo int, hi int, k int) int {
	dp := map[int]int{}
	n := hi - lo + 1
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = lo + i
	}
	sort.Slice(nums, func(i, j int) bool {
		a := resolve(dp, nums[i])
		b := resolve(dp, nums[j])
		if a < b {
			return true
		} else if a == b {
			return nums[i] < nums[j]
		}
		return false
	})
	return nums[k-1]
}
