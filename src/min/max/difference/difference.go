package difference

func minMaxDifference(num int) int {
	digits := make([]int, 0, 9)
	for num != 0 {
		digits = append(digits, num%10)
		num /= 10
	}
	n := len(digits)
	base := 1
	for i := 1; i < n; i++ {
		base *= 10
	}
	a, b := -1, -1
	var num1, num2 int
	for i := n - 1; i >= 0; i-- {
		if a != -1 {
			if digits[i] == a {
				num1 += base * 9
			} else {
				num1 += base * digits[i]
			}
		} else {
			if digits[i] != 9 {
				a = digits[i]
			}
			num1 += base * 9
		}

		if b != -1 {
			if digits[i] != b {
				num2 += base * digits[i]
			}
		} else {
			if digits[i] != 0 {
				b = digits[i]
			}
		}
		base /= 10
	}
	return num1 - num2
}
