package drop

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func superEggDrop(k int, n int) int {
	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, n+1)
	}
	return dfs(dp, k, n)
}

func dfs(dp [][]int, k int, n int) int {
	if n <= 1 || k == 1 {
		return n
	}
	if dp[k][n] != 0 {
		return dp[k][n]
	}
	l := 1
	r := n
	for l <= r {
		mid := (l + r) / 2
		// 把蛋扔在mid层后，破与未破两种情况对应的蛋的数量和未确定层数
		// 为了取最坏的情况，所以取最大值
		r1 := dfs(dp, k-1, mid-1) + 1
		r2 := dfs(dp, k, n-mid) + 1
		temp := max(r1, r2)
		if dp[k][n] == 0 || temp < dp[k][n] {
			dp[k][n] = temp
		}
		if r1 >= r2 {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return dp[k][n]
}
