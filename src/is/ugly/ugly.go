package ugly

func isUgly(n int) bool {
	if n == 0 {
		return false
	}
	bases := []int{2, 3, 5}
	for _, base := range bases {
		for n != 1 {
			if n%base == 0 {
				n = n / base
			} else {
				break
			}
		}
		if n == 1 {
			break
		}
	}
	return n == 1
}
