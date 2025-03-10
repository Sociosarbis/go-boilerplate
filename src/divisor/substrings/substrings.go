package substrings

import "fmt"

func divisorSubstrings(num int, k int) int {
	base := 1
	for k != 0 {
		base *= 10
		k--
	}
	var ret int
	divisor := num
	for divisor != 0 {
		fmt.Println(divisor, base)
		d := divisor % base
		if d != 0 {
			if num%d == 0 {
				ret++
			}
		}
		if divisor < base {
			break
		}
		divisor /= 10
	}
	return ret
}
