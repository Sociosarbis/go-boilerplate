package number

func maximum69Number(num int) int {
	var incr int
	base := 1
	temp := num
	for temp != 0 {
		digit := temp % 10
		if digit == 6 {
			incr = 3 * base
		}
		temp /= 10
		base *= 10
	}
	return num + incr
}
