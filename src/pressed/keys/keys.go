package keys

const mod int = 1e9 + 7

func countTexts(pressedKeys string) int {
	dp := [2][4]int{}
	var index int
	var prevIndex int
	var k byte
	var options int
	dp[0][0] = 1
	n := len(pressedKeys)
	for i := 1; i < n; i++ {
		k = pressedKeys[i]
		index = i & 1
		prevIndex = 1 - index
		dp[index][0] = dp[prevIndex][0]
		for j := 1; j < 4; j++ {
			dp[index][0] = (dp[index][0] + dp[prevIndex][j]) % mod
			dp[index][j] = 0
		}
		if k == pressedKeys[i-1] {
			if k == '9' || k == '7' {
				options = 4
			} else {
				options = 3
			}
			for j := 1; j < options; j++ {
				dp[index][j] = dp[prevIndex][j-1]
			}
		}
	}
	index = (n - 1) & 1
	ret := dp[index][0]
	for i := 1; i < 4; i++ {
		ret = (ret + dp[index][i]) % mod
	}
	return ret
}
