package flips3

func minFlips(s string) int {
	n := len(s)
	if n == 1 {
		return 0
	}
	dp := [2][]int{
		make([]int, n),
		make([]int, n),
	}
	for i := 0; i <= 1; i++ {
		prev := byte(i) + '0'
		if s[0] != prev {
			dp[i][0] = 1
		}
		for j := 1; j < n; j++ {
			if s[j] == prev {
				dp[i][j] = dp[i][j-1] + 1
			} else {
				dp[i][j] = dp[i][j-1]
			}
			prev = '1' - prev + '0'
		}
	}
	var out int
	if dp[0][n-1] < dp[1][n-1] {
		out = dp[0][n-1]
	} else {
		out = dp[1][n-1]
	}
	for i := 0; i <= 1; i++ {
		prev := byte(i) + '0'
		var temp int
		if s[n-1] != prev {
			temp = 1
		}
		if temp+dp[1-i][n-2] < out {
			out = temp + dp[1-i][n-2]
		}
		for j := n - 2; j > 0; j-- {
			if s[j] == prev {
				temp++
			}
			if temp+dp[1-i][j-1] < out {
				out = temp + dp[1-i][j-1]
			}
			prev = '1' - prev + '0'
		}
	}
	return out
}
