package change

import "sort"

func change(amount int, coins []int) int {
	sort.Ints(coins)
	dp := [5001]int{}
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i += 1 {
			if dp[i-coin] != 0 {
				dp[i] += dp[i-coin]
			}
		}
	}
	return dp[amount]
}
