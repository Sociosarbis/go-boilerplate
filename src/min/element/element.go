package element

func sum(num int) int {
	var out int
	for num != 0 {
		out += num % 10
		num /= 10
	}
	return out
}

func minElement(nums []int) int {
	out := 36
	for _, num := range nums {
		res := sum(num)
		if res < out {
			out = res
		}
	}
	return out
}
