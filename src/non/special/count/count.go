package count

import "math"

func isPrime(num int) bool {
	if num == 1 {
		return false
	}
	end := int(math.Sqrt(float64(num)))
	for i := 2; i <= end; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func nonSpecialCount(l int, r int) int {
	total := r - l + 1
	start := int(math.Sqrt(float64(l)))
	if start*start < l {
		start++
	}
	end := int(math.Sqrt(float64(r)))
	if end*end > r {
		end--
	}
	for i := start; i <= end; i++ {
		if isPrime(i) {
			total--
		}
	}
	return total
}
