package strings

func validStrings(n int) []string {
	ret := [][]byte{{'0'}, {'1'}}
	for i := 1; i < n; i++ {
		temp := make([][]byte, 0, len(ret))
		for _, item := range ret {
			if item[i-1] == '0' {
				temp = append(temp, append(item, '1'))
			} else {
				t := make([]byte, 0, i)
				copy(t, item)
				temp = append(temp, append(t, '0'), append(item, '1'))
			}
		}
		ret = temp
	}
	out := make([]string, 0, len(ret))
	for _, item := range ret {
		out = append(out, string(item))
	}
	return out
}
