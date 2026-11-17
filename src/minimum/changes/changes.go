package changes

import "fmt"

func getFactors(n int) []int {
	out := make([]int, 0, 15)
	out = append(out, 1)
	i := 2
	for ; i*i < n; i++ {
		if n%i == 0 {
			out = append(out, i, n/i)
		}
	}
	if i*i == n {
		out = append(out, i)
	}
	return out
}

func getCost(s string, l, r, d int) int {
	var out int
	n := (r - l + 1) / d
	for i := d - 1; i >= 0; i-- {
		j, k := 0, n-1
		for j < k {
			if s[l+i+d*j] != s[l+i+d*k] {
				out++
			}
			j++
			k--
		}
	}
	if l == 0 && n == 10 {
		fmt.Println(out)
	}
	return out
}

func minimumChanges(s string, k int) int {
	n := len(s)
	factors := make([][]int, n+1)
	factors[1] = []int{1}
	for i := 2; i <= n; i++ {
		factors[i] = getFactors(i)
	}
	cost := make([][]int, n)
	for i := 0; i < n; i++ {
		cost[i] = make([]int, n)
		for j := i + 1; j < n; j++ {
			cost[i][j] = j - i
			for _, d := range factors[j-i+1] {
				c := getCost(s, i, j, d)
				if c < cost[i][j] {
					cost[i][j] = c
				}
			}
		}
	}
	dp := make([][2]int, n)
	for i := n - 1; i >= 0; i-- {
		dp[i][0] = cost[i][n-1]
	}
	for i := 1; i < k; i++ {
		index := i & 1
		prevIndex := 1 - index
		for j := n - (i+1)*2; j >= 0; j-- {
			dp[j][index] = n

			for r := j + 2; r+i*2 <= n; r++ {
				if cost[j][r-1]+dp[r][prevIndex] < dp[j][index] {
					dp[j][index] = cost[j][r-1] + dp[r][prevIndex]
				}
			}
		}
	}
	index := (k - 1) & 1
	return dp[0][index]
}
