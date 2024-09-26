package sum

func getDigitSum(num int) int {
	var ret int
	for num != 0 {
		ret += num % 10
		num /= 10
	}
	return ret
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func differenceOfSum(nums []int) int {
	var ret int
	for _, num := range nums {
		ret += abs(num - getDigitSum(num))
	}
	return ret
}
