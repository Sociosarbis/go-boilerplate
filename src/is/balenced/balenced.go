package balenced

func isBalanced(num string) bool {
	var ret int
	for i, n := range num {
		if i&1 == 0 {
			ret -= int(n - 48)
		} else {
			ret += int(n - 48)
		}
	}
	return ret == 0
}
