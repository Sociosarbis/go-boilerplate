package permutation

const mask int = 1e9 + 7

func numberOfPermutations(n int, requirements [][]int) int {
	rs := make(map[int]int, len(requirements))
	// 题目输入限定了n个元素排列的最多逆序值
	var max int
	for _, r := range requirements {
		rs[r[0]] = r[1]
		if r[1] > max {
			max = r[1]
		}
	}
	dp := [2][]int{make([]int, max+1), make([]int, max+1)}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		index := i % 2
		prevIndex := 1 - index
		var start int
		end := max
		if c, ok := rs[i-1]; ok {
			start = c
			if end > start {
				end = start
			}
			dp[index] = make([]int, max+1)
		}
		for j := start; j <= end; j++ {
			var temp int
			var e int
			if j-i+1 > e {
				e = j - i + 1
			}
			// 因为i比之前的元素大，所以可以增加0到i-1个逆序
			for k := j; k >= e; k-- {
				temp = (temp + dp[prevIndex][k]) % mask
			}
			dp[index][j] = temp
		}
	}
	index := n % 2
	return dp[index][max]
}
