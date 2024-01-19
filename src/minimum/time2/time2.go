package time2

import (
	"sort"

	"github.com/boilerplate/src/slice"
)

func minimumTime(nums1 []int, nums2 []int, x int) int {
	pairs := make([][2]int, len(nums1))
	sum2 := 0
	for i, num := range nums1 {
		pairs[i] = [2]int{nums2[i], num}
		x -= num
		sum2 += nums2[i]
	}
	if x >= 0 {
		return 0
	}
	n := len(pairs)

	sort.Sort(slice.NewSortable(
		pairs,
		func(a, b [2]int) bool {
			return a[0] < b[0]
		}))

	/** 前i个执行j次操作后的最大值 */
	dp := make([][]int, n+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	ret := -1

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			temp := dp[i-1][j-1] + pairs[i-1][0]*j + pairs[i-1][1]
			if temp > dp[i-1][j] {
				dp[i][j] = temp
			} else {
				dp[i][j] = dp[i-1][j]
			}
			if (ret == -1 || j < ret) && (sum2*j-dp[i][j] <= x) {
				ret = j
			}
		}
	}
	return ret
}
