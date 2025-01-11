package key

func generateKey(num1 int, num2 int, num3 int) int {
	var ret int
	base := 1
	for i := 0; i < 4; i++ {
		a, b, c := num1%10, num2%10, num3%10
		if a > b {
			a = b
		}
		if a > c {
			a = c
		}
		if a != 0 {
			ret += base * a
		}
		num1, num2, num3 = num1/10, num2/10, num3/10
		base *= 10
	}
	return ret
}
