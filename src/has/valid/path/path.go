package path

type empty struct{}

func inherit(src map[int]empty, d int) map[int]empty {
	return extend(make(map[int]empty, len(src)), src, d)

}

func extend(target, src map[int]empty, d int) map[int]empty {
	for c := range src {
		if c+d >= 0 {
			target[c+d] = empty{}
		}
	}
	return target
}

func hasValidPath(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])
	if grid[0][0] != '(' || grid[m-1][n-1] != ')' {
		return false
	}
	dp := make([][]map[int]empty, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]map[int]empty, n)
	}
	dp[0][0] = map[int]empty{
		1: {},
	}
	for i := 1; i < n; i++ {
		var d int
		if grid[0][i] == '(' {
			d = 1
		} else {
			d = -1
		}
		dp[0][i] = inherit(dp[0][i-1], d)

	}
	for i := 1; i < m; i++ {
		index := i & 1
		prevIndex := 1 - index
		var d int
		if grid[i][0] == '(' {
			d = 1
		} else {
			d = -1
		}
		dp[index][0] = inherit(dp[prevIndex][0], d)
		for j := 1; j < n; j++ {
			if grid[i][j] == '(' {
				d = 1
			} else {
				d = -1
			}
			dp[index][j] = extend(inherit(dp[prevIndex][j], d), dp[index][j-1], d)
		}
	}
	index := (m - 1) & 1
	if _, ok := dp[index][n-1][0]; ok {
		return true
	}
	return false
}
