package bit

func evenOddBit(n int) []int {
	var i int
	ret := make([]int, 2)
	for n != 0 {
		if n&1 == 1 {
			ret[i]++
		}
		i = 1 - i
		n >>= 1
	}
	return ret
}
