package of2

func reorderedPowerOf2(n int) bool {
	counter := make([]int, 10)
	for n > 0 {
		counter[n%10] += 1
		n /= 10
	}
	counterCopy := make([]int, 10)
	copy(counterCopy, counter)
	for i := 0; i < 31; i += 1 {
		num := 1 << i
		for num > 0 {
			j := num % 10
			counter[j] -= 1
			if counter[j] < 0 {
				break
			}
			num /= 10
		}
		ret := true
		for _, count := range counter {
			if count != 0 {
				ret = false
				break
			}
		}
		if ret {
			return ret
		}
		for j := range counter {
			counter[j] = counterCopy[j]
		}
	}
	return false
}
