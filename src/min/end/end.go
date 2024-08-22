package end

func minEnd(n int, x int) int64 {
	ret := int64(x)
	i := 0
	n -= 1
	for n != 0 {
		for ret&(1<<i) != 0 {
			i++
		}
		ret |= int64(n&1) << i
		i++
		n >>= 1
	}
	return ret
}
