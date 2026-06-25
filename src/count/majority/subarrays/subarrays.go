package subarrays

func countMajoritySubarrays(nums []int, target int) int {
	n := len(nums)
	prefixSum := make([]int, n+1)
	for i, num := range nums {
		if num == target {
			prefixSum[i+1] = prefixSum[i] + 1
		} else {
			prefixSum[i+1] = prefixSum[i]
		}
	}
	var out int
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if (prefixSum[j+1]-prefixSum[i])*2 > j-i+1 {
				out++
			}
		}
	}
	return out
}
