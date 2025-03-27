package cost3

func minimumCost(s string) int64 {
	n := len(s)
	suffix := make([]int64, n+1)
	for i := n - 2; i >= 0; i-- {
		if s[i] == s[i+1] {
			suffix[i] = suffix[i+1]
		} else {
			suffix[i] = suffix[i+1] + int64(n-i-1)
		}
	}
	ret := suffix[0]
	var prefix int64
	for i := 1; i < n; i++ {
		if s[i] != s[i-1] {
			prefix += int64(i)
		}
		if prefix+suffix[i] < ret {
			ret = prefix + suffix[i]
		}
	}
	return ret
}
