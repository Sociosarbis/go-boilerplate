package subarrays

func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
	n := len(nums)
	prefixSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		if nums[i]%modulo == k {
			prefixSum[i+1] = prefixSum[i] + 1
		} else {
			prefixSum[i+1] = prefixSum[i]
		}
	}
	n = len(prefixSum)
	// count[j] = (count[i] + modulo - k) % modulo
	m := map[int]int{}
	var ret int64
	for i := 0; i < n; i++ {
		if c, ok := m[(prefixSum[i]+modulo-k)%modulo]; ok {
			ret += int64(c)
		}
		key := prefixSum[i] % modulo
		if c, ok := m[key]; ok {
			m[key] = c + 1
		} else {
			m[key] = 1
		}
	}
	return ret
}
