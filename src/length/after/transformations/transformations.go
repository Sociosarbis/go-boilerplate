package transformations

const mask int = 1e9 + 7

func lengthAfterTransformations(s string, t int) int {
	n := t + 26
	dp := make([][]int, 26)
	for i := 0; i < 26; i++ {
		dp[i] = make([]int, 26)
	}
	dp[0][0] = 1
	for i := 1; i < n; i++ {
		prevIndex := (i - 1) % 26
		index := i % 26
		copy(dp[index][1:], dp[prevIndex][:])
		dp[index][0] = 0
		last := dp[prevIndex][25]
		if last != 0 {
			dp[index][0] += last
			dp[index][1] += last
			dp[index][0] %= mask
			dp[index][1] %= mask
		}
	}
	var ret int
	for _, c := range s {
		index := (int(c-97) + t) % 26
		res := dp[index]
		for i := 0; i < 26; i++ {
			ret = (ret + res[i]) % mask
		}
	}
	return ret
}
