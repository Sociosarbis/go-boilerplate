package palindrome

import (
	"math"
	"strconv"
)

func largestPalindrome(n int) int {
	if n > 1 {
		upper := int(math.Pow(10, float64(n))) - 1
		lower := upper / 10

		temp := make([]byte, n*2)
		for i := upper; i > lower; i-- {
			tempNum := i
			for j := n - 1; j >= 0; j-- {
				digit := tempNum % 10
				tempNum /= 10
				temp[j] = byte(digit + 48)
				temp[len(temp)-1-j] = temp[j]
			}
			if p, err := strconv.ParseInt(string(temp), 10, 64); err == nil {
				for j := int64(upper); j*j > p; j-- {
					if p%j == 0 {
						return int(p % 1337)
					}
				}
			}

		}
	}
	return 9
}
