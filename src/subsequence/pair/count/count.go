package count

const mask int = 1e9 + 7

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a < b {
		a, b = b, a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func subsequencePairCount(nums []int) int {
	dp := [2][][]int{}
	for i := 0; i < 2; i++ {
		dp[i] = make([][]int, 201)
		for j := 0; j <= 200; j++ {
			dp[i][j] = make([]int, 201)
		}
	}
	dp[1][0][0] = 1
	for i, num := range nums {
		index := i & 1
		prevIndex := 1 - index
		for j := 0; j <= 200; j++ {
			for k := 0; k <= 200; k++ {
				dp[index][j][k] = dp[prevIndex][j][k]
			}
		}
		for j := 0; j <= 200; j++ {
			for k := 0; k <= 200; k++ {
				if dp[prevIndex][j][k] != 0 {
					g1, g2 := gcd(j, num), gcd(k, num)
					dp[index][g1][k] = (dp[index][g1][k] + dp[prevIndex][j][k]) % mask
					dp[index][j][g2] = (dp[index][j][g2] + dp[prevIndex][j][k]) % mask
				}
			}
		}
	}
	index := (len(nums) - 1) & 1
	var out int
	for i := 1; i <= 200; i++ {
		out = (out + dp[index][i][i]) % mask
	}
	return out
}
