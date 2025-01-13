package array

func waysToSplitArray(nums []int) int {
	var total int64
	for _, num := range nums {
		total += int64(num)
	}
	var a int64
	var ret int
	n := len(nums) - 1
	for i := 0; i < n; i++ {
		a += int64(nums[i])
		if a >= total-a {
			ret++
		}
	}
	return ret
}
