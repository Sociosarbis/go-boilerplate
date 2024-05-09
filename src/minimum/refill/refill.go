package refill

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	l := 0
	r := len(plants) - 1
	a := capacityA
	b := capacityB
	var ret int
	for l <= r {
		if l != r {
			if plants[l] > a {
				ret++
				a = capacityA - plants[l]
			} else {
				a -= plants[l]
			}

			if plants[r] > b {
				ret++
				b = capacityB - plants[r]
			} else {
				b -= plants[r]
			}
		} else {
			if a < b {
				a = b
			}
			if plants[l] > a {
				ret++
			}
		}
		l++
		r--
	}
	return ret
}
