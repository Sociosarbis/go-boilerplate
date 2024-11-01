package boost

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func maxEnergyBoost(energyDrinkA []int, energyDrinkB []int) int64 {
	dp := [2][2]int64{}
	dp[0][0], dp[0][1] = int64(energyDrinkA[0]), int64(energyDrinkB[0])
	n := len(energyDrinkA)
	for i := 1; i < n; i++ {
		index := i % 2
		prevIndex := 1 - index
		dp[index][0] = max(dp[prevIndex][0]+int64(energyDrinkA[i]), dp[prevIndex][1])
		dp[index][1] = max(dp[prevIndex][1]+int64(energyDrinkB[i]), dp[prevIndex][0])
	}
	index := (n - 1) % 2
	return max(dp[index][0], dp[index][1])
}
