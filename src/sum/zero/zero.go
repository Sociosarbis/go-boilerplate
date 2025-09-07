package zero

func sumZero(n int) []int {
	out := make([]int, 0, n)
	from := n / 2
	for i := -from; i < 0; i++ {
		out = append(out, i)
	}
	for i := 1; i <= from; i++ {
		out = append(out, i)
	}
	if n%2 == 1 {
		out = append(out, 0)
	}
	return out
}
