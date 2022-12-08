package delivering

func boxDelivering(boxes [][]int, portsCount int, maxBoxes int, maxWeight int) int {
	n := len(boxes)
	// g[i] 存储转移方程所需的中间变量，表示从第i个箱子从仓库开始一次运输时的dp[i]（i - 1的最优值） - neg[i + 1]
	// 因为从第i个箱子开始，所以i与之前的箱子是否不同目标港口，不需要考虑，所以是i+1
	g := make([]int, n)
	// dp[i] 运完第i - 1个箱子所需的最小步数
	dp := make([]int, n+1)
	// sum[i] i前的所有箱子的重量加上第一个箱子
	sum := make([]int, n+1)
	// neg[i] i前有多少次目标港口的转换
	neg := make([]int, n+1)

	sum[0] = boxes[0][1]
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + boxes[i-1][1]
		if i > 1 {
			if boxes[i-2][0] == boxes[i-1][0] {
				neg[i] = neg[i-1]
			} else {
				neg[i] = neg[i-1] + 1
			}
		}
	}
	q := []int{0}
	for i := 1; i <= n; i++ {
		for len(q) != 0 && (i-q[0] > maxBoxes || sum[i]-sum[q[0]] > maxWeight) {
			q = q[1:]
		}
		// 一次有效的从仓库出发再回到仓库的运输，等于不同目标港口的箱子种类数加2
		dp[i] = g[q[0]] + neg[i] + 2
		if i < n {
			g[i] = dp[i] - neg[i+1]
			// 如果不比以第i个箱子开始更优，则出栈
			for len(q) != 0 && g[i] <= g[q[len(q)-1]] {
				q = q[:len(q)-1]
			}
			q = append(q, i)
		}
	}
	return dp[n]
}
