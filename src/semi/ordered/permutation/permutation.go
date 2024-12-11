package permutation

func semiOrderedPermutation(nums []int) int {
	n := len(nums)
	var visited bool
	var ret int
	for i, num := range nums {
		switch num {
		case 1:
			if visited {
				return ret + i
			} else {
				ret = i
				visited = true
			}
		case n:
			if visited {
				return ret + n - 1 - i
			} else {
				ret = n - 2 - i
			}
		}
	}
	return ret
}
