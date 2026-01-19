package length

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxSideLength(mat [][]int, threshold int) int {
	m, n := len(mat), len(mat[0])
	prefixSum := make([][]int, m)
	for i := 0; i < m; i++ {
		prefixSum[i] = make([]int, n)
	}
	prefixSum[0][0] = mat[0][0]
	for i := 1; i < n; i++ {
		prefixSum[0][i] = prefixSum[0][i-1] + mat[0][i]
	}
	for i := 1; i < m; i++ {
		prefixSum[i][0] = prefixSum[i-1][0] + mat[i][0]
		for j := 1; j < n; j++ {
			prefixSum[i][j] = prefixSum[i][j-1] + prefixSum[i-1][j] - prefixSum[i-1][j-1] + +mat[i][j]
		}
	}
	var out int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := min(m-i, n-j); k > out; k-- {
				var temp int
				if i > 0 {
					temp += prefixSum[i-1][j+k-1]
				}
				if j > 0 {
					temp += prefixSum[i+k-1][j-1]
				}
				if i > 0 && j > 0 {
					temp -= prefixSum[i-1][j-1]
				}
				temp = prefixSum[i+k-1][j+k-1] - temp
				if temp <= threshold {
					out = k
					break
				}
			}
		}
	}
	return out
}
