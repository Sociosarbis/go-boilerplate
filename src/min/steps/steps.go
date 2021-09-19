package steps

func minSteps(n int) int {
	primes := []int{1}
	dp := make([]int, n+1)
	dp[1] = 0
	for i := 2; i <= n; i += 1 {
		for j := len(primes) - 1; j >= 0; j -= 1 {
			if i%primes[j] == 0 {
				if j == 0 {
					primes = append(primes, i)
					dp[i] = i
				} else {
					dp[i] = dp[i/primes[j]] + primes[j]
				}
				break
			}
		}
	}
	return dp[n]
}
