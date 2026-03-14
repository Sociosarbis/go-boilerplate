package string

func dfs(s *[]byte, n int, k *int) {
	var j byte = 'a'
	for ; j < 'd'; j++ {
		if len(*s) == 0 || j != (*s)[len(*s)-1] {
			*s = append(*s, j)
			if len(*s) == n {
				*k--
			} else {
				dfs(s, n, k)
			}
			if *k == 0 {
				return
			}
			*s = (*s)[:len(*s)-1]
		}
	}
}

func getHappyString(n int, k int) string {
	s := make([]byte, 0, n)
	dfs(&s, n, &k)
	return string(s)
}
