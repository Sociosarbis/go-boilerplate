package value

import "sort"

func kthLargestValue(matrix [][]int, k int) int {
	m := len(matrix)
	n := len(matrix[0])
	dp := make([]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			index := i*n + j
			if i != 0 {
				top := (i-1)*n + j
				if j == 0 {
					dp[index] = matrix[i][j] ^ dp[top]
				} else {
					left := i*n + j - 1
					topLeft := top - 1
					dp[index] = dp[left] ^ dp[top] ^ dp[topLeft] ^ matrix[i][j]
				}
			} else {
				if j != 0 {
					dp[index] = matrix[i][j] ^ dp[index-1]
				} else {
					dp[0] = matrix[0][0]
				}
			}
		}
	}
	sort.Ints(dp)
	return dp[len(dp)-k]
}
