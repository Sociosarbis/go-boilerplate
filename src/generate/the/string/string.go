package string

func generateTheString(n int) string {
	c := byte('a')
	ret := make([]byte, n)
	i := 0
	if n&1 == 0 {
		ret[i] = c
		c++
		i++
	}
	for ; i < n; i++ {
		ret[i] = c
	}
	return string(ret)
}
