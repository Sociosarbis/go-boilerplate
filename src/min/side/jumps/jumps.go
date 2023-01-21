package jumps

func minSideJumps(obstacles []int) int {
	n := len(obstacles)
	dp := make([][]int, 2)

	for i := range dp {
		dp[i] = make([]int, 3)
	}

	max := 3 * n

	dp[0][0] = 1
	dp[0][2] = 1

	var index int
	for i := 1; i < n; i++ {
		index = i & 1
		for j := 0; j < 3; j++ {
			if j+1 != obstacles[i] {
				temp := max
				for k := 0; k < 3; k++ {
					if dp[1-index][k] != -1 {
						var d int
						if k != j {
							if obstacles[i] == k+1 {
								continue
							} else {
								d = 1
							}
						}
						v := dp[1-index][k] + d
						if v < temp {
							temp = v
						}
					}
				}
				dp[index][j] = temp
			} else {
				dp[index][j] = -1
			}
		}
	}
	ret := max
	for i := 0; i < 3; i++ {
		if dp[index][i] != -1 && dp[index][i] < ret {
			ret = dp[index][i]
		}
	}
	return ret
}
