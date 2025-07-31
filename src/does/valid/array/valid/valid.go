package valid

func doesValidArrayExist(derived []int) bool {
	n := len(derived)
	var temp int
	for i := 1; i < n; i++ {
		temp = derived[i-1] ^ temp
	}
	if temp == derived[n-1] {
		return true
	}
	temp = 1
	for i := 1; i < n; i++ {
		temp = derived[i-1] ^ temp
	}
	return temp^1 == derived[n-1]
}
