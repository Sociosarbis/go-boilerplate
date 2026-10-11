package squares

func sumOfSquares(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0] * nums[0]
	}
	out := nums[0]*nums[0] + nums[n-1]*nums[n-1]
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			out += nums[i-1] * nums[i-1]
			b := n / i
			if b != i {
				out += nums[b-1] * nums[b-1]
			}
		}
	}
	return out
}
