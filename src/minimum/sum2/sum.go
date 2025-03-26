package sum2

func minimumSum(n int, k int) int {
	a := k / 2
	if a >= n {
		return (1 + n) * n / 2
	}
	return ((1+a)*a + (k*2+n-a-1)*(n-a)) / 2
}
