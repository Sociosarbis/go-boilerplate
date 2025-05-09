package permutations

const MOD int64 = 1e9 + 7

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func countBalancedPermutations(num string) int {
	var total int
	n := len(num)
	counter := [10]int{}
	for i := 0; i < n; i++ {
		d := num[i] - '0'
		counter[d]++
		total += int(d)
	}
	if total%2 != 0 {
		return 0
	}
	target := total / 2
	maxOdd := (n + 1) / 2
	comb := make([][]int64, maxOdd+1)
	for i := 0; i <= maxOdd; i++ {
		comb[i] = make([]int64, maxOdd+1)
		comb[i][0] = 1
		comb[i][i] = comb[i][0]
		for j := 1; j < i; j++ {
			// 这里应该是假设有黑白两种球，一共有i个球，求j个白球时的组合数
			// 前i-1个球可以放了j个白球，或者放了j-1个白球，
			// 因为前i-1个球确定了，那下一个球该放什么也是确定的
			comb[i][j] = (comb[i-1][j] + comb[i-1][j-1]) % MOD
		}
	}
	f := make([][]int64, target+1)
	for i := 0; i <= target; i++ {
		f[i] = make([]int64, maxOdd+1)
	}
	f[0][0] = 1
	/** 数字数量，数字和 */
	var psum, totSum int
	for i := 0; i <= 9; i++ {
		psum += counter[i]
		totSum += i * counter[i]
		/** 遍历可能的奇数位个数，同时保证偶数位不能大于可能的最大值n - maxOdd */
		for oddCount := min(psum, maxOdd); oddCount >= max(0, psum-(n-maxOdd)); oddCount-- {
			evenCount := psum - oddCount
			/** 遍历各个可能的奇数和，同时保证偶数和不能大于target */
			for curr := min(totSum, target); curr >= max(0, totSum-target); curr-- {
				var res int64
				/** j个i放在奇数位，最多可以放evenCount个i到偶数位，j个i的和不能大于总的奇数位和curr */
				for j := max(0, counter[i]-evenCount); j <= min(counter[i], oddCount) && i*j <= curr; j++ {
					// 在oddCount个奇数位下放j个i，在eventCount个偶数位下放counter[i]-j个i，因为是互相独立的所以乘起来
					ways := comb[oddCount][j] * comb[evenCount][counter[i]-j] % MOD
					// 其他的位置就放之前的数字，curr和odd都是随着i的增加而增加，所以每次遍历i的时候，f要不被重置为0，要不重新计算
					// 而且只有f不是0的时候表示包含了之前数字的在curr-i*j,oddCount-j条件下的所有可能
					res = (res + ways*f[curr-i*j][oddCount-j]%MOD) % MOD
				}
				f[curr][oddCount] = res % MOD
			}
		}
	}
	return int(f[target][maxOdd])
}
