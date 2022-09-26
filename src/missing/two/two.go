package two

func missingTwo(nums []int) []int {
	n := len(nums) + 2
	sumSquare := 0
	for i := 1; i <= n; i++ {
		sumSquare += i * i
	}
	sum := (1 + n) * n / 2

	for _, num := range nums {
		sum -= num
		sumSquare -= num * num
	}

	for i := 1; i <= n; i++ {
		if i*i+(sum-i)*(sum-i) == sumSquare {
			return []int{i, sum - i}
		}
	}
	return make([]int, 0)
}
