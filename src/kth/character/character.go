package character

func kthCharacter(k int) byte {
	dp := make([]byte, k)
	dp[0] = 'a'
	var i int
	for i+1 < k {
		for j := 0; j <= i && i+1+j < k; j++ {
			dp[i+1+j] = (dp[j]-'a'+1)%26 + 'a'
		}
		i = (i+1)*2 - 1
	}
	return dp[k-1]
}
