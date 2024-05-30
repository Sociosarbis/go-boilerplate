package length2

func maximumLength(s string) int {
	n := len(s)
	dp := [26][]int{}
	for i := 0; i < 26; i++ {
		dp[i] = make([]int, 0, 3)
	}
	var i int
	for i < n {
		c := s[i]
		j := i + 1
		for ; j < n && s[j] == c; j++ {
		}
		l := j - i
		end := l - 3
		if end < 0 {
			end = 0
		}
		for j := l; j > end; j-- {
			var k int
			index := c - 97
			n2 := len(dp[index])
			for ; k < n2; k++ {
				if j > dp[index][k] {
					break
				}
			}
			if k < 3 {
				if k >= n2 {
					dp[index] = append(dp[index], j)
					n2++
				} else {
					copy(dp[index][k+1:], dp[index][k:])
					dp[index][k] = j
				}
			} else {
				break
			}
		}
		i = j
	}
	var ret int
	for i := 0; i < 26; i++ {
		if len(dp[i]) == 3 && dp[i][2] > ret {
			ret = dp[i][2]
		}
	}
	if ret == 0 {
		return -1
	}
	return ret
}
