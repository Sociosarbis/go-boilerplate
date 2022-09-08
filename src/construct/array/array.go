package array

func constructArray(n int, k int) []int {
	ret := make([]int, n)
	l := 1
	r := k + 1
	for i := 0; i < len(ret); i++ {
		if l == r+1 {
			l = k + 2
		}
		if i&1 == 0 || l > r {
			ret[i] = l
			l++
		} else {
			ret[i] = r
			r--
		}
	}
	return ret
}
