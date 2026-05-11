package digits

func add(nums []int, num int) []int {
	s := len(nums)
	for num != 0 {
		nums = append(nums, num%10)
		num /= 10
	}
	for e := len(nums) - 1; e > s; e-- {
		nums[e], nums[s] = nums[s], nums[e]
		s++
	}
	return nums
}

func separateDigits(nums []int) []int {
	n := len(nums)
	out := make([]int, 0, n*5)
	for _, num := range nums {
		out = add(out, num)
	}
	return out
}
