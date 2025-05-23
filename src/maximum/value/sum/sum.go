package sum

func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	n := len(nums)
	g := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	dp := make([][2]int64, n)
	dfs(g, dp, nums, 0, -1, k)
	return dp[0][0]
}

func dfs(g [][]int, dp [][2]int64, nums []int, c, p, k int) {
	if len(g[c]) == 1 && p != -1 {
		dp[c][0], dp[c][1] = int64(nums[c]), int64(nums[c]^k)
	} else {
		for _, child := range g[c] {
			if child != p {
				dfs(g, dp, nums, child, c, k)
			}
		}
		var minDiff int64 = 1e9
		var maxDiff int64 = -1e9
		var temp int64
		var count int
		n := len(g[c])
		if p != -1 {
			n--
		}
		for _, child := range g[c] {
			if child != p {
				diff := dp[child][1] - dp[child][0]
				if dp[child][1] > dp[child][0] {
					count++
					temp += dp[child][1]
					if diff < minDiff {
						minDiff = diff
					}
				} else {
					temp += dp[child][0]
					if diff > maxDiff {
						maxDiff = diff
					}
				}
			}
		}
		a, b := int64(nums[c]), int64(nums[c]^k)
		if count&1 == 0 {
			a, b = b, a
		}
		if p != -1 {
			// 利用只有一个父节点的特性，改变通向父节点的路径
			dp[c][1] = temp + a
			if count > 0 && temp-minDiff+b > dp[c][1] {
				dp[c][1] = temp - minDiff + b
			}
			if count+1 <= n {
				if temp+maxDiff+b > dp[c][1] {
					dp[c][1] = temp + maxDiff + b
				}
			}
		}
		dp[c][0] = temp + b
		if temp-minDiff+a > dp[c][0] {
			dp[c][0] = temp - minDiff + a
		}
		if count+1 <= n {
			if temp+maxDiff+a > dp[c][0] {
				dp[c][0] = temp + maxDiff + a
			}
		}
	}
}
