package distance5

func maxDistance(colors []int) int {
	var out int
	n := len(colors)
	for j := n - 1; j >= 0; j-- {
		if colors[j] != colors[0] {
			out = j
			break
		}
	}
	for i := 0; i < n; i++ {
		if colors[i] != colors[n-1] {
			if n-1-i > out {
				out = n - 1 - i
			}
			break
		}
	}
	return out
}
