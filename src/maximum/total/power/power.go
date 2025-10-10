package power

import "sort"

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func maximumTotalDamage(power []int) int64 {
	n := len(power)
	counter := make(map[int]int, n)
	for _, power := range power {
		if c, ok := counter[power]; ok {
			counter[power] = c + 1
		} else {
			counter[power] = 1
		}
	}
	var i int
	for k := range counter {
		power[i] = k
		i++
	}
	n = len(counter)
	power = power[:n]
	sort.Ints(power)
	dp := make([][2]int64, n)
	for i := 0; i < n; i++ {
		var a int64
		for j := i - 1; j >= 0; j-- {
			if power[j]+2 < power[i] {
				a = max(a, max(dp[j][0], dp[j][1]))
				break
			}
		}
		dp[i][1] = a + int64(power[i])*int64(counter[power[i]])
		a = 0
		for j := i - 1; j >= 0; j-- {
			if power[j]+2 < power[i] {
				break
			}
			a = max(a, max(dp[j][0], dp[j][1]))
		}
		dp[i][0] = a
	}
	return max(dp[n-1][0], dp[n-1][1])
}
