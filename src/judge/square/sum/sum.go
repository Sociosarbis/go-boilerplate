package sum

import "math"

func judgeSquareSum(c int) bool {
	n := int(math.Sqrt(float64(c)))
	if n*n == c {
		return true
	}
	n = int(math.Sqrt(float64(c) / 2))
	for i := 1; i <= n; i++ {
		remain := c - i*i
		b := int(math.Sqrt(float64(remain)))
		if b*b == remain {
			return true
		}
	}
	return false
}
