package distance

import "sort"

const max int64 = 2 * 1e11

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func minimumTotalDistance(robot []int, factory [][]int) int64 {
	sort.Ints(robot)
	sort.Slice(factory, func(i, j int) bool {
		return factory[i][0] < factory[j][0]
	})
	n := len(robot)
	m := len(factory)
	dp := make([][]int64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int64, m)
		for j := 0; j < m; j++ {
			dp[i][j] = max
		}
	}
	if factory[0][1] != 0 {
		dp[0][0] = int64(abs(robot[0] - factory[0][0]))
	}
	for j := 1; j < m; j++ {
		dp[0][j] = dp[0][j-1]
		if factory[j][1] != 0 {
			temp := int64(abs(robot[0] - factory[j][0]))
			if temp < dp[0][j] {
				dp[0][j] = temp
			}
		}
	}

	for i := 1; i < n; i++ {
		if i < factory[0][1] {
			dp[i][0] = dp[i-1][0] + int64(abs(robot[i]-factory[0][0]))
		}
		for j := 1; j < m; j++ {
			dp[i][j] = dp[i][j-1]
			k := factory[j][1]
			if k > i+1 {
				k = i + 1
			}

			var cost int64
			for diff := 0; diff < k; diff++ {
				from := i - diff
				var temp int64
				if from > 0 {
					temp = dp[from-1][j-1]
				}
				cost += int64(abs(robot[from] - factory[j][0]))
				if temp+cost < dp[i][j] {
					dp[i][j] = temp + cost
				}
			}
		}
	}
	return dp[n-1][m-1]
}
