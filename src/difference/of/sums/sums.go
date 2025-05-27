package sums

func differenceOfSums(n int, m int) int {
	var ret int
	for i := 1; i <= n; i++ {
		if i%m == 0 {
			ret -= i
		} else {
			ret += i
		}
	}
	return ret
}
