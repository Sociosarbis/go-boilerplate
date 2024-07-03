package number

func sumOfTheDigitsOfHarshadNumber(x int) int {
	temp := x
	var ret int
	for temp != 0 {
		ret += temp % 10
		temp /= 10
	}
	if x%ret == 0 {
		return ret
	}
	return -1
}
