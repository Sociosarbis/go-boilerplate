package rotations

func count(tops, bottoms []int, num, i int) int {
	n := len(tops)
	var ret int
	if i == 0 {
		for j := 1; j < n; j++ {
			if tops[j] != num {
				if bottoms[j] == num {
					ret++
				} else {
					return -1
				}
			}
		}
	} else {
		for j := 1; j < n; j++ {
			if bottoms[j] != num {
				if tops[j] == num {
					ret++
				} else {
					return -1
				}
			}
		}
	}
	return ret
}

func minDominoRotations(tops []int, bottoms []int) int {
	ret := -1
	if temp := count(tops, bottoms, tops[0], 0); temp != -1 && (temp < ret || ret == -1) {
		ret = temp
	}
	if temp := count(tops, bottoms, bottoms[0], 0); temp != -1 && (temp+1 < ret || ret == -1) {
		ret = temp + 1
	}
	if temp := count(tops, bottoms, tops[0], 1); temp != -1 && (temp+1 < ret || ret == -1) {
		ret = temp + 1
	}
	if temp := count(tops, bottoms, bottoms[0], 1); temp != -1 && (temp < ret || ret == -1) {
		ret = temp
	}
	return ret
}
