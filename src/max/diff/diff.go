package diff

func maxDiff(num int) int {
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
	var bFactor int
	var num1, num2 int
	for i := n - 1; i >= 0; i-- {
		if a == -1 {
			if digits[i] != 9 {
				a = digits[i]
			}
			num1 += base * 9
		} else {
			if digits[i] == a {
				num1 += base * 9
			} else {
				num1 += base * digits[i]
			}
		}

		if b == -1 {
			if i+1 == n {
				if digits[i] != 1 {
					b = digits[i]
					bFactor = 1
				}
				num2 += base
			} else {
				if digits[i] > 1 {
					b = digits[i]
				} else if digits[i] != 0 {
					num2 += base * digits[i]
				}
			}
		} else {
			if digits[i] == b {
				if bFactor != 0 {
					num2 += base
				}
			} else {
				num2 += base * digits[i]
			}
		}

		base /= 10
	}
	return num1 - num2
}
