package order

func reverseDegree(s string) int {
	var out int
	for i, c := range s {
		out += (i + 1) * int(123-c)
	}
	return out
}
