package array3

func resultArray(nums []int, k int) []int64 {
	dp := [2][]int64{
		make([]int64, k),
		make([]int64, k),
	}
	empty := make([]int64, k)
	out := make([]int64, k)
	index := 1
	for _, num := range nums {
		prevIndex := index
		index = 1 - index
		copy(dp[index], empty)
		ni := num % k
		out[ni]++
		dp[index][ni] = 1
		for i := 0; i < k; i++ {
			if dp[prevIndex][i] != 0 {
				ni = (i * num) % k
				count := dp[prevIndex][i]
				dp[index][ni] += count
				out[ni] += count
			}
		}
	}
	return out
}
