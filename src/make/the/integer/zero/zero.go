package zero

func makeTheIntegerZero(num1 int, num2 int) int {
	if num2 >= num1 {
		return -1
	}
	i := 1
	for i <= 32 {
		diff := num1 - num2*i
		if diff < i {
			break
		}
		var count int
		for diff != 0 {
			diff &= diff - 1
			count++
		}
		if count <= i {
			return i
		}
		i++
	}
	return -1
}
