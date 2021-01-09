package profit

func maxProfit(prices []int) int {
	dp := [2][2]int{{-prices[0], 0}, {-prices[0], 0}}
	for i := 1; i < len(prices); i++ {
		if -prices[i] > dp[0][0] {
			dp[0][0] = -prices[i]
		}

		if dp[0][0]+prices[i] > dp[0][1] {
			dp[0][1] = dp[0][0] + prices[i]
		}

		if dp[0][1]-prices[i] > dp[1][0] {
			dp[1][0] = dp[0][1] - prices[i]
		}

		if dp[1][0]+prices[i] > dp[1][1] {
			dp[1][1] = dp[1][0] + prices[i]
		}
	}

	if dp[1][1] > dp[0][1] {
		return dp[1][1]
	}
	return dp[0][1]
}
