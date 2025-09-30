package sum

func triangularSum(nums []int) int {
	n := len(nums)
	cur, next := nums, make([]int, n-1)
	for n > 1 {
		for i := n - 2; i >= 0; i-- {
			next[i] = (cur[i] + cur[i+1]) % 10
		}
		cur, next = next, cur
		n--
	}
	if n&1 == 1 {
		return cur[0]
	}
	return next[0]
}
