package drop

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func twoEggDrop(n int) int {
	// dp[i] 等于确定i层所需要的次数
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = i
		// 当第一个鸡蛋在第j层落下时
		// 如果破裂，则需要从第一层开始逐层测试，总共需要j次
		// 如果不破裂，则相当于要从j+1到i中确定，等价于要确定i-j层，再加上第一个鸡蛋投下的这次
		for j := 1; j < i; j++ {
			dp[i] = min(dp[i], max(j, dp[i-j]+1))
		}
	}
	return dp[n]
}
