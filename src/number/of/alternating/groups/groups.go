package groups

func numberOfAlternatingGroups(colors []int) int {
	var ret int
	n := len(colors)
	i1 := n - 1
	for i, color := range colors {
		i2 := (i + 1) % n
		if color != colors[i1] && color != colors[i2] {
			ret++
		}
		i1 = i
	}
	return ret
}
