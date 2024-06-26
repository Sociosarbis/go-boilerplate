package perm

const mod int = 1e9 + 7
const baseMask int = (1 << 14) - 1

func specialPerm(nums []int) int {
	n := len(nums)

	dp := map[int]int{}
	for i := 0; i < n; i++ {
		dp[encode(0, i)] = 1
	}
	for i := 2; i <= n; i++ {
		nextDp := map[int]int{}
		for j := 0; j < n; j++ {
			num := nums[j]
			for mask, c := range dp {
				if mask&(1<<j) == 0 {
					index := decode(mask)
					if nums[index]%num == 0 || num%nums[index] == 0 {
						mask = encode(mask, j)
						if count, ok := nextDp[mask]; ok {
							nextDp[mask] = (count + c) % mod
						} else {
							nextDp[mask] = c
						}
					}
				}
			}
		}
		dp = nextDp
	}
	var ret int
	for _, count := range dp {
		ret = (ret + count) % mod
	}
	return ret
}

func encode(mask int, i int) int {
	return ((mask | (1 << i)) & baseMask) | ((i + 1) << 14)
}

func decode(mask int) int {
	return (mask >> 14) - 1
}
