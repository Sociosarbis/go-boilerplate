package stickers

func minStickers(stickers []string, target string) int {
	n := len(target)
	m := 1 << n
	dp := make([]int, m)
	for i := 1; i < m; i++ {
		dp[i] = n + 1
	}
	for i := 0; i < m; i++ {
		if dp[i] != n+1 {
			for _, sticker := range stickers {
				cur := i
				for _, c := range sticker {
					for j := 0; j < n; j++ {
						if target[j] == byte(c) && ((cur>>j)&1 == 0) {
							cur |= 1 << j
							break
						}
					}
				}
				if dp[i]+1 < dp[cur] {
					dp[cur] = dp[i] + 1
				}
			}
		}
	}
	if dp[m-1] != n+1 {
		return dp[m-1]
	}
	return -1
}
