package array

func constructTransformedArray(nums []int) []int {
	n := len(nums)
	out := make([]int, n)
	for i, num := range nums {
		if num > 0 {
			out[i] = nums[(i+num)%n]
		} else if num < 0 {
			out[i] = nums[(i+num+100*n)%n]
		} else {
			out[i] = num
		}
	}
	return out
}
