package changes

func countKeyChanges(s string) int {
	n := len(s)
	var ret int
	for i := 1; i < n; i++ {
		v := int(s[i]) - int(s[i-1])
		if !(v == 0 || v == 32 || v == -32) {
			ret++
		}
	}
	return ret
}
