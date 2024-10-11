package pairs

import "math"

func numberOfPairs(nums1 []int, nums2 []int, k int) int64 {
	m := make(map[int]int, 2000)
	var ret int64
	for _, num := range nums1 {
		if num%k == 0 {
			v := num / k
			n := int(math.Sqrt(float64(v)))
			for i := 1; i <= n; i++ {
				if v%i == 0 {
					j := v / i
					if c, ok := m[i]; ok {
						m[i] = c + 1
					} else {
						m[i] = 1
					}
					if i != j {
						if c, ok := m[j]; ok {
							m[j] = c + 1
						} else {
							m[j] = 1
						}
					}
				}
			}
		}
	}
	for _, num := range nums2 {
		if c, ok := m[num]; ok {
			ret += int64(c)
		}
	}
	return ret
}
