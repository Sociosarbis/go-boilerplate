package array

import (
	"sort"
)

func findOriginalArray(changed []int) []int {
	n := len(changed)
	if n&1 == 1 {
		return []int{}
	}
	sort.Ints(changed)
	counter := make(map[int]int, n)

	for _, num := range changed {
		if num&1 == 0 {
			if count, ok := counter[num]; ok {
				counter[num] = count + 1
			} else {
				counter[num] = 1
			}
		}
	}

	ret := make([]int, 0, n>>1)

	for _, num := range changed {
		if num&1 == 0 {
			if count, ok := counter[num]; ok && count > 0 {
				counter[num] = count - 1
			} else {
				continue
			}
		}
		doubled := num << 1
		if count, ok := counter[doubled]; ok && count > 0 {
			counter[doubled] = count - 1
		} else {
			return []int{}
		}
		ret = append(ret, num)
		if len(ret)<<1 == n {
			break
		}
	}

	return ret
}
