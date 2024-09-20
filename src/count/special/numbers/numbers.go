package numbers

func countSpecialNumbers(n int) int {
	digits := make([]int, 0, 10)
	for n != 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	return count(digits, len(digits)-1, 0, false) - 1
}

func count(digits []int, i, mask int, free bool) int {
	if i < 0 {
		return 1
	}
	if free {
		if mask == 0 {
			var ret int
			for index := 1; index < 10; index++ {
				ret += count(digits, i-1, mask|(1<<index), free)
			}
			ret += count(digits, i-1, mask, free)
			return ret
		}
		var used int
		for index := 0; index < 10; index++ {
			if mask&(1<<index) != 0 {
				used++
			}
		}
		n := i + 1
		ret := 1
		option_count := 10 - used
		for index := 0; index < n; index++ {
			ret *= option_count - index
		}
		return ret
	}
	var ret int
	for num := digits[i]; num >= 0; num-- {
		if mask&(1<<num) == 0 {
			var nextMask int
			if !(num == 0 && mask == 0) {
				nextMask = mask | (1 << num)
			}
			ret += count(digits, i-1, nextMask, num != digits[i])
		}
	}
	return ret
}
