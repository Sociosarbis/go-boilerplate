package collection

func garbageCollection(garbage []string, travel []int) int {
	trucks := [3]int{0, 0, 0}
	gToT := map[rune]int{
		77: 0,
		80: 1,
		71: 2,
	}

	var ret int
	for i, g := range garbage {
		for _, c := range g {
			t := gToT[c]
			j := trucks[t]
			if j < i {
				for ; j < i; j++ {
					ret += travel[j]
				}
				trucks[t] = i
			}
			ret++
		}
	}
	return ret
}
