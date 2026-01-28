package cost7

import "sort"

func min(a, b int) int {
	if a == -1 {
		return b
	}
	if b == -1 {
		return a
	}
	if a < b {
		return a
	}
	return b
}

func minCost(grid [][]int, k int) int {
	dp := [2][][]int{}
	m, n := len(grid), len(grid[0])
	for i := 0; i < 2; i++ {
		dp[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = make([]int, n)
		}
	}
	for i := 1; i < n; i++ {
		dp[0][0][i] = dp[0][0][i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		dp[0][i][0] = dp[0][i-1][0] + grid[i][0]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[0][i][j] = min(dp[0][i-1][j], dp[0][i][j-1]) + grid[i][j]
		}
	}
	out := dp[0][m-1][n-1]
	points := make([][2]int, 0, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			points = append(points, [2]int{i, j})
		}
	}
	sort.Slice(points, func(i, j int) bool {
		return grid[points[i][0]][points[i][1]] < grid[points[j][0]][points[j][1]]
	})
	for t := 1; t <= k; t++ {
		index := t & 1
		prevIndex := 1 - index
		minCost := -1
		i := len(points) - 1
		j := i
		// 先以传送的方式更新初始值
		for ; i >= 0; i-- {
			minCost = min(minCost, dp[prevIndex][points[i][0]][points[i][1]])
			if i > 0 && grid[points[i][0]][points[i][1]] == grid[points[i-1][0]][points[i-1][1]] {
				continue
			}
			for k := i; k <= j; k++ {
				dp[index][points[k][0]][points[k][1]] = minCost
			}
			j = i - 1
		}
		for i := 1; i < n; i++ {
			if dp[index][0][i-1] != -1 {
				dp[index][0][i] = min(dp[index][0][i], dp[index][0][i-1]+grid[0][i])
			}
		}
		for i := 1; i < m; i++ {
			if dp[index][i-1][0] != -1 {
				dp[index][i][0] = min(dp[index][i][0], dp[index][i-1][0]+grid[i][0])
			}
		}
		for i := 1; i < m; i++ {
			for j := 1; j < n; j++ {
				prev := min(dp[index][i-1][j], dp[index][i][j-1])
				if prev != -1 {
					dp[index][i][j] = min(dp[index][i][j], prev+grid[i][j])
				}
			}
		}
		out = min(dp[index][m-1][n-1], out)
	}
	return out
}
