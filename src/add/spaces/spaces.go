package spaces

func addSpaces(s string, spaces []int) string {
	n1 := len(s)
	n2 := len(spaces)
	n := n1 + n2
	ret := make([]byte, 0, n)
	var i int
	var j int
	for i < n1 {
		if j < n2 && i == spaces[j] {
			ret = append(ret, ' ')
			j++
		} else {
			ret = append(ret, s[i])
			i++
		}
	}
	return string(ret)
}
