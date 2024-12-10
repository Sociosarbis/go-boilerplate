package dialer

var DIRS = [8][2]int{{-2, -1}, {-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {-1, -2}, {1, -2}}

const m int = 1e9 + 7

func knightDialer(n int) int {
	dp := [2][12]int{}
	for i := 0; i < 9; i++ {
		dp[0][i] = 1
	}
	dp[0][10] = 1
	for i := 1; i < n; i++ {
		index := i % 2
		prevIndex := 1 - index
		dp[index] = [12]int{}
		for j := 0; j < 4; j++ {
			for k := 0; k < 3; k++ {
				ii := j*3 + k
				prevValue := dp[prevIndex][ii]
				if prevValue != 0 {
					for _, dir := range DIRS {
						jj, kk := j+dir[0], k+dir[1]
						if jj >= 0 && jj < 4 && kk >= 0 && kk < 3 && (jj != 3 || kk == 1) {
							iii := jj*3 + kk
							dp[index][iii] = (dp[index][iii] + prevValue) % m
						}
					}
				}
			}
		}
	}
	index := (n - 1) % 2
	var ret int
	for i := 0; i < 12; i++ {
		ret = (ret + dp[index][i]) % m
	}
	return ret
}
