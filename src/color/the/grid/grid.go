package grid

const mask int = 1e9 + 7
const offset int = 3

func colorTheGrid(m int, n int) int {
	masks := []int{1, 2, 4}
	for i := 1; i < m; i++ {
		nextMasks := []int{}
		for j := 0; j < offset; j++ {
			for _, m := range masks {
				if m&((1<<j)<<(offset*(i-1))) == 0 {
					nextMasks = append(nextMasks, m|((1<<j)<<(offset*i)))
				}
			}
		}
		masks = nextMasks
	}
	dp := [2]map[int]int{}
	dp[0], dp[1] = make(map[int]int, len(masks)), make(map[int]int, len(masks))
	for _, mask := range masks {
		dp[0][mask] = 1
	}
	// 遍历列
	for i := 1; i < n; i++ {
		index := i & 1
		prevIndex := 1 - index
		for _, a := range masks {
			dp[index][a] = 0
			for _, b := range masks {
				if a&b == 0 {
					dp[index][a] = (dp[index][a] + dp[prevIndex][b]) % mask
				}
			}
		}
	}
	index := (n - 1) & 1
	var ret int
	for _, m := range masks {
		ret = (ret + dp[index][m]) % mask
	}
	return ret
}
