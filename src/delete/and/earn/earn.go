package earn

import (
	"sort"
)

func deleteAndEarn(nums []int) int {
	earns := [1000]int{}
	options := []int{}
	for _, num := range nums {
		if earns[num-1] == 0 {
			options = append(options, num)
		}
		earns[num-1] += num
	}
	sort.Ints(options)
	dp := make([][2]int, len(options))
	for i := 0; i < len(options); i++ {
		option := options[i] - 1
		if i != 0 {
			if dp[i-1][0] > dp[i-1][1] {
				dp[i][0] = dp[i-1][0]
			} else {
				dp[i][0] = dp[i-1][1]
			}
			if options[i-1] == option {
				dp[i][1] = dp[i-1][0] + earns[option]
			} else {
				dp[i][1] = dp[i][0] + earns[option]
			}
		} else {
			dp[i][0] = 0
			dp[i][1] = earns[option]
		}
	}
	if dp[len(dp)-1][0] > dp[len(dp)-1][1] {
		return dp[len(dp)-1][0]
	} else {
		return dp[len(dp)-1][1]
	}
}
