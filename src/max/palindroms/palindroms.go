package palindroms

func maxPalindromes(s string, k int) int {
	n := len(s)
	if k == 1 {
		return n
	}
	dp := make([]int, n)
	start := (k - 1) / 2
	end := n - start
	for i := start; i < end; i++ {
		if i > 0 {
			if dp[i] < dp[i-1] {
				dp[i] = dp[i-1]
			}
		}
		l := i
		r := i
		for l-1 >= 0 && r+1 < n {
			if s[l-1] == s[r+1] {
				l--
				r++
				if r-l+1 >= k {
					var temp int
					if l > 0 {
						temp = dp[l-1] + 1
					} else {
						temp = 1
					}
					if dp[r] < temp {
						dp[r] = temp
					}
					break
				}
			} else {
				break
			}
		}
		l = i
		r = i + 1
		for l >= 0 && r < n {
			if s[l] == s[r] {
				if r-l+1 >= k {
					var temp int
					if l > 0 {
						temp = dp[l-1] + 1
					} else {
						temp = 1
					}
					if dp[r] < temp {
						dp[r] = temp
					}
					break
				}
				l--
				r++
			} else {
				break
			}
		}
	}
	out := dp[end-1]
	for i := end; i < n; i++ {
		if out < dp[i] {
			out = dp[i]
		}
	}
	return out
}
