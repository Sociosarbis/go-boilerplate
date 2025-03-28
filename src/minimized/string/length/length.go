package length

func minimizedStringLength(s string) int {
	var mask int
	for _, c := range s {
		mask |= 1 << (c - 97)
	}
	var ret int
	for mask != 0 {
		ret++
		mask &= mask - 1
	}
	return ret
}
