package array

func resultsArray(nums []int, k int) []int {
	j := 0
	n := len(nums)
	ret := make([]int, n-k+1)
	n = len(ret)
	for i := 0; i < n; i++ {
		if j < i {
			j = i
		}
		for j-i+1 < k && nums[j+1] == nums[j]+1 {
			j++
		}
		if j-i+1 == k {
			ret[i] = nums[j]
		} else {
			ret[i] = -1
		}
	}
	return ret
}
