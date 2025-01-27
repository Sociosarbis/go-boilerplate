package jump

func jump(nums []int) int {
	n := len(nums)
	dp := make([]int, 0, n)
	dp = append(dp, 0)
	var i int
	for j, num := range nums {
		for j < dp[i] {
			i++
		}
		next := j + num
		if next > dp[i] {
			if i+1 >= len(dp) {
				dp = append(dp, next)
			} else if next > dp[i+1] {
				dp[i+1] = next
				dp = dp[:i+2]
			}
		}
	}
	ret := len(dp) - 1
	for ; ret > 0; ret-- {
		if dp[ret-1]+1 < n {
			break
		}
	}
	return ret
}
