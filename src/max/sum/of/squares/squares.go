package squares

func maxSumOfSquares(num int, sum int) string {
	out := make([]byte, 0, num)
	for num != 0 {
		if sum > 9 {
			out = append(out, '9')
			sum -= 9
		} else if sum != 0 {
			out = append(out, '0'+byte(sum))
			sum = 0
		} else {
			out = append(out, '0')
		}
		num--
	}
	if sum != 0 {
		return ""
	}
	return string(out)
}
