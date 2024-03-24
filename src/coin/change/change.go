package change

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	filteredCoins := make([]int, 0, len(coins))
	for _, coin := range coins {
		if coin < amount {
			dp[coin] = 1
			filteredCoins = append(filteredCoins, coin)
		} else if coin == amount {
			return 1
		}
	}
	for i := 1; i <= amount; i++ {
		if dp[i] != 0 {
			continue
		}
		for _, coin := range filteredCoins {
			if coin < i {
				v := dp[i-coin]
				if v != 0 {
					temp := v + 1
					if dp[i] == 0 || temp < dp[i] {
						dp[i] = temp
					}
				}
			}
		}
	}
	if dp[amount] == 0 {
		return -1
	}
	return dp[amount]
}
