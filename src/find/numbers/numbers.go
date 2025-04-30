package numbers

func isEven(n int) bool {
	var ret int
	for n != 0 {
		n /= 10
		ret++
	}
	return ret&1 == 0
}

func findNumbers(nums []int) int {
	var ret int
	for _, num := range nums {
		if isEven(num) {
			ret++
		}
	}
	return ret
}
