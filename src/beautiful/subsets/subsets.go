package subsets

func beautifulSubsets(nums []int, k int) int {
	n := len(nums)
	dp := make([]bool, 1<<n)
	dp[0] = true
	var ret int
	for i, num := range nums {
		var mask int
		for j := 0; j < i; j++ {
			if num+k == nums[j] || num-k == nums[j] {
				mask |= 1 << j
			}
		}
		end := 1 << i
		for j := 0; j < end; j++ {
			if dp[j] && j&mask == 0 {
				ret++
				dp[j|end] = true
			}
		}
	}
	return ret
}
