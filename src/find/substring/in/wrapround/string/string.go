package string

func findSubstringInWraproundString(p string) int {
	dp := make([]int, 26)

	l := 0
	r := 0
	for l < len(p) {
		r = l + 1
		for r < len(p) {
			if p[r]-1 == p[r-1] || (p[r-1]-96)%26 == p[r]-97 {
				r++
			} else {
				break
			}
		}
		for i := l; i < r; i++ {
			if r-i > dp[p[i]-97] {
				dp[p[i]-97] = r - i
			}
		}
		l = r
	}

	ret := 0

	for _, count := range dp {
		if count > 0 {
			ret += count
		}
	}

	return ret
}
