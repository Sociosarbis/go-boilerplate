package swap

import (
	"sort"
)

func maximumSwap(num int) int {
	if num == 0 {
		return 0
	}
	digits := []int{}
	for num != 0 {
		digits = append(digits, num%10)
		num /= 10
	}
	temp := make([]int, len(digits))
	copy(temp, digits)
	sort.Ints(temp)
	i := len(digits) - 1
	for i >= 0 {
		if temp[i] != digits[i] {
			newVal := temp[i]
			oldVal := digits[i]
			digits[i] = newVal
			i = 0
			for i < len(temp) {
				if digits[i] == newVal {
					digits[i] = oldVal
					break
				} else {
					i++
				}
			}
			break
		} else {
			i--
		}
	}
	base := 1
	ret := 0
	for i := 0; i < len(digits); i++ {
		ret += digits[i] * base
		base *= 10
	}
	return ret
}
