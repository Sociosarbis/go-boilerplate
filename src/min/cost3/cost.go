package cost3

import "sort"

func minCost(n int, cuts []int) int {
	sort.Ints(cuts)
	m := len(cuts) + 2
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, m)
	}
	max := len(cuts) * n

	for i := 2; i < m; i++ {
		for j := 0; j+i < m; j++ {
			temp := max
			for k := j + 1; k < j+i; k++ {
				t := dp[j][k] + dp[k][j+i]
				var left int
				if j > 0 {
					left = cuts[j-1]
				}
				var right int
				if j+i+1 >= m {
					right = n
				} else {
					right = cuts[j+i-1]
				}
				if t+right-left < temp {
					temp = t + right - left
				}
			}
			dp[j][j+i] = temp
		}
	}
	return dp[0][m-1]
}
