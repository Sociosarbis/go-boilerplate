package index

func sumDigits(num int) int {
	var out int
	for num != 0 {
		out += num % 10
		num /= 10
	}
	return out
}

func smallestIndex(nums []int) int {
	for i, num := range nums {
		if sumDigits(num) == i {
			return i
		}
	}
	return -1
}
