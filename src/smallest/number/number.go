package number

func smallestNumber(n int) int {
	var count int
	for n != 0 {
		n >>= 1
		count++
	}
	return (1 << count) - 1
}
