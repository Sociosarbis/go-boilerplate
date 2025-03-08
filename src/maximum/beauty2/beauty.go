package beauty2

import "sort"

func getPartialValue(flowers []int, prefix []int64, newFlowers int64, minI, target, partial int) int64 {
	n := len(flowers)
	r := sort.Search(len(flowers), func(i int) bool {
		if i < minI {
			return false
		}
		if flowers[i] >= target {
			return false
		}
		has := prefix[n] - prefix[i]
		need := int64(flowers[i]) * int64(n-i)
		if need <= has {
			return true
		}
		if newFlowers >= need-has {
			return true
		}
		return false
	})
	if r != n {
		v := ((newFlowers + prefix[n] - prefix[r]) / int64(n-r))
		if r != minI {
			target = flowers[r-1]
		}
		if v >= int64(target) {
			v = int64(target - 1)
		}
		return v * int64(partial)
	}
	return 0
}

func maximumBeauty(flowers []int, newFlowers int64, target int, full int, partial int) int64 {
	sort.Slice(flowers, func(i, j int) bool {
		return flowers[i] > flowers[j]
	})
	n := len(flowers)
	prefix := make([]int64, n+1)
	for i, flower := range flowers {
		if flower > target {
			flower = target
		}
		prefix[i+1] = prefix[i] + int64(flower)
	}
	var i int
	for ; i < n; i++ {
		if flowers[i] < target {
			break
		}
	}

	ret := int64(full)*int64(i) + getPartialValue(flowers, prefix, newFlowers, i, target, partial)
	i++
	for ; i <= n; i++ {
		has := prefix[i]
		need := int64(target) * int64(i)
		if has+newFlowers < need {
			break
		}
		remain := newFlowers - (need - has)
		v := int64(i)*int64(full) + getPartialValue(flowers, prefix, remain, i, target, partial)
		if v > ret {
			ret = v
		}
	}
	return ret
}
