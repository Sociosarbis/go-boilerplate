package string2

func getSmallestString(s string) string {
	n := len(s)
	ret := []byte(s)
	for i := 1; i < n; i++ {
		if ret[i-1] > ret[i] && ret[i]%2 == ret[i-1]%2 {
			ret[i-1], ret[i] = ret[i], ret[i-1]
			break
		}
	}
	return string(ret)
}
