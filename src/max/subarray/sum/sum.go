package sum

func maxSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	prefixSum := make([]int64, k)
	var temp int64
	for i := 0; i < k; i++ {
		temp += int64(nums[i])
		prefixSum[i] = temp
	}
	ret := temp
	if prefixSum[k-1] > 0 {
		prefixSum[k-1] = 0
	}
	for i := k; i < n; i++ {
		temp += int64(nums[i])
		if temp-prefixSum[i%k] > ret {
			ret = temp - prefixSum[i%k]
		}
		if temp < prefixSum[i%k] {
			prefixSum[i%k] = temp
		}
	}
	return ret
}
