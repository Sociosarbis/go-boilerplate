package length

func maximumLength(s string) int {
	n := len(s)
	visited := make(map[string]int, n)
	var i int
	for i < n {
		c := s[i]
		j := i + 1
		for ; j < n; j++ {
			if s[j] != c {
				break
			}
		}
		end := j - i
		for l := j - i; l > 0; l-- {
			temp := s[i : i+l]
			delta := end - l + 1
			if count, ok := visited[temp]; ok {
				visited[temp] = count + delta
			} else {
				visited[temp] = delta
			}
		}
		i = j
	}
	ret := -1
	for k, v := range visited {
		if v >= 3 {
			if len(k) > ret {
				ret = len(k)
			}
		}
	}
	return ret
}
