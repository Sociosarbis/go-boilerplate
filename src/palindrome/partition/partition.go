package partition

func count(s string) int {
	var l int
	r := len(s) - 1
	var ret int
	for l < r {
		if s[l] != s[r] {
			ret++
		}
	}
	return ret
}

func palindromePartition(s string, k int) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k+1)
		end := k
		if end > i+1 {
			end = i + 1
		}
		for j := 2; j <= end; j++ {
			dp[i][j] = n
		}
	}
	for i := 1; i < n; i++ {
		dp[i][1] = count(s[:i+1])
		for j := 0; j < i; j++ {
			end := k - 1
			if end > j+1 {
				end = j + 1
			}
			c := count(s[j+1 : i+1])
			for ii := 1; ii <= end; ii++ {
				temp := dp[j][ii] + c
				if dp[i][ii+1] > temp {
					dp[i][ii+1] = temp
				}
			}
		}
	}
	return dp[n-1][k]
}
