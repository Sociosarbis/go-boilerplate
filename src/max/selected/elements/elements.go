package elements

import "sort"

func maxSelectedElements(nums []int) int {
	sort.Ints(nums)
	out := 1
	dp := [2][2]int{}
	n := len(nums)
	dp[0][0], dp[0][1] = 1, 1
	for i := 1; i < n; i++ {
		index := i & 1
		prevIndex := 1 - index
		if nums[i] == nums[i-1] {
			dp[index][0], dp[index][1] = dp[prevIndex][0], dp[prevIndex][1]
		} else {
			dp[index][0], dp[index][1] = 1, 1
		}
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				if nums[i-1]+k+1 == nums[i]+j && dp[prevIndex][k]+1 > dp[index][j] {
					dp[index][j] = dp[prevIndex][k] + 1
				}
			}
			if dp[index][j] > out {
				out = dp[index][j]
			}
		}
	}
	return out
}
