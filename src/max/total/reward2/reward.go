package reward2

import "sort"

const FULL_MASK = (1 << 30) - 1

func splitMask(mask, offset int) (int, int, int) {
	remain := offset % 30
	return offset / 30, (mask << remain) & FULL_MASK, mask >> (30 - remain)
}

func maxTotalReward(rewardValues []int) int {
	sort.Ints(rewardValues)
	n := len(rewardValues)
	max := rewardValues[n-1]
	dp := make([]int, max*2/30+1)
	dp[0] = 1
	for _, v := range rewardValues {
		for j := 0; j*30 < v; j++ {
			m := dp[j]
			if (j+1)*30 > v {
				m = dp[j] & ((1 << (v % 30)) - 1)
			}
			i, m1, m2 := splitMask(m, v)
			dp[i+j] |= m1
			if m2 != 0 {
				dp[i+j+1] |= m2
			}
		}
	}
	n = len(dp)
	for i := n - 1; i >= 0; i-- {
		if dp[i] != 0 {
			for j := 29; j >= 0; j-- {
				if dp[i]&(1<<j) != 0 {
					return i*30 + j
				}
			}
		}
	}
	return 0
}
