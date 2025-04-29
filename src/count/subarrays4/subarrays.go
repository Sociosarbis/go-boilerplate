package subarrays4

func countSubarrays(nums []int, k int) int64 {
	var max int
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	indices := []int{}
	for i, num := range nums {
		if num == max {
			indices = append(indices, i)
		}
	}
	var start int
	n := len(nums)
	end := len(indices) - k
	var ret int64
	for i := 0; i <= end; i++ {
		ret += int64(n-indices[i+k-1]) * int64(indices[i]-start+1)
		start = indices[i] + 1
	}
	return ret
}
