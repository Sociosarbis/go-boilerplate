package doubled

import (
	"sort"
)

func canReorderDoubled(arr []int) bool {
	pos := []int{}
	neg := []int{}
	counter := map[int]int{}
	zeroes := 0
	for _, num := range arr {
		counter[num] += 1
		if num > 0 {
			pos = append(pos, num)
		} else if num < 0 {
			neg = append(neg, num)
		} else {
			zeroes += 1
		}
	}

	if len(pos)&1 != 0 {
		return false
	} else if len(neg)&1 != 0 {
		return false
	} else if zeroes&1 != 0 {
		return false
	}

	sort.Ints(pos)
	step := len(pos) / 2
	for i := 0; i < len(pos); i += 1 {
		if step == 0 {
			break
		}
		if counter[pos[i]] != 0 {
			if counter[2*pos[i]] == 0 {
				return false
			}
			counter[pos[i]] -= 1
			counter[2*pos[i]] -= 1
			step -= 1
		}
	}
	if step != 0 {
		return false
	}
	sort.Ints(neg)
	step = len(neg) / 2
	for i := len(neg) - 1; i >= 0; i -= 1 {
		if step == 0 {
			break
		}
		if counter[neg[i]] != 0 {
			if counter[2*neg[i]] == 0 {
				return false
			}
			counter[neg[i]] -= 1
			counter[neg[i]*2] -= 1
			step -= 1
		}
	}
	if step != 0 {
		return false
	}
	return true
}
