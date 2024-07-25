package operations

func minimumOperations(num string) int {
	n := len(num)
	indices := [2]int{n, n}
	for i := n - 1; i >= 0; i-- {
		if indices[0] == n {
			if num[i] == '5' {
				indices[0] = i
			}
		} else {
			if num[i] == '7' || num[i] == '2' {
				return n - i - 2
			}
		}
		if indices[1] == n {
			if num[i] == '0' {
				indices[1] = i
			}
		} else {
			if num[i] == '0' || num[i] == '5' {
				return n - i - 2
			}
		}
	}
	if indices[1] != n {
		return n - 1
	}
	return n
}
