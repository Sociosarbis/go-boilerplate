package numbers

func findMinFibonacciNumbers(k int) int {
	pair := []int{1, 1}
	mem := []int{1}
	i := 1
	for pair[i] < k {
		pair[1-i] = pair[0] + pair[1]
		mem = append(mem, pair[1-i])
		i = 1 - i
	}
	ret := 0
	i = len(mem) - 1
	for k != 0 {
		for mem[i] > k {
			i -= 1
		}
		k -= mem[i]
		ret += 1
	}
	return ret
}
