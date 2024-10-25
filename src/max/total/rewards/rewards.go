package rewards

import "sort"

func maxTotalReward(rewardValues []int) int {
	sort.Ints(rewardValues)
	n := len(rewardValues)
	max := rewardValues[n-1]
	dp := make([]bool, max*2)
	dp[0] = true
	for _, v := range rewardValues {
		for i := 0; i < v; i++ {
			if dp[i] && !dp[i+v] {
				dp[i+v] = true
			}
		}
	}
	var ret int
	n = len(dp)
	for i := n - 1; i >= 0; i-- {
		if dp[i] {
			ret = i
			break
		}
	}
	return ret
}
