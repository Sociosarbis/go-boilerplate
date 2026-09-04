package index

func firstStableIndex(nums []int, k int) int {
	n := len(nums)
	min := make([]int, n)
	min[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] < min[i+1] {
			min[i] = nums[i]
		} else {
			min[i] = min[i+1]
		}
	}
	var max int
	for i, num := range nums {
		if num > max {
			max = num
		}
		if max-min[i] <= k {
			return i
		}
	}
	return -1
}
