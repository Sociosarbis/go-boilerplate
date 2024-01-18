package removal

import "sort"

func minimumRemoval(beans []int) int64 {
	sort.Ints(beans)
	var sum int64
	for _, bean := range beans {
		sum += int64(bean)
	}
	ret := sum
	for i, bean := range beans {
		if i == 0 || beans[i-1] != bean {
			temp := sum - int64(len(beans)-i)*int64(bean)
			if temp < ret {
				ret = temp
			}
		}
	}
	return ret
}
