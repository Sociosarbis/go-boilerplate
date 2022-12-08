package operation3

func minOperations(nums1 []int, nums2 []int) int {
	a := make([]int, 6)
	sumA := 0
	b := make([]int, 6)
	sumB := 0
	for _, num := range nums1 {
		sumA += num
		a[num-1]++
	}
	for _, num := range nums2 {
		sumB += num
		b[num-1]++
	}
	if sumA == sumB {
		return 0
	}
	if sumB > sumA {
		a, b = b, a
		sumA, sumB = sumB, sumA
	}
	i := 5
	j := 0
	ret := 0
	for i > 0 && a[i] == 0 {
		i--
	}
	for j < 5 && b[j] == 0 {
		j++
	}
	for i > 0 || j < 5 {
		diff := sumA - sumB
		if i >= 5-j {
			d1 := i * a[i]
			if diff <= d1 {
				if diff%i == 0 {
					return ret + diff/i
				} else {
					return ret + diff/i + 1
				}
			} else {
				ret += a[i]
			}
			sumA -= d1
			a[i] = 0
			for i > 0 && a[i] == 0 {
				i--
			}
		} else {
			d2 := (5 - j) * b[j]
			if diff <= d2 {
				if diff%(5-j) == 0 {
					return ret + diff/(5-j)
				} else {
					return ret + diff/(5-j) + 1
				}
			} else {
				ret += b[j]
			}
			sumB += d2
			b[j] = 0
			for j < 5 && b[j] == 0 {
				j++
			}
		}
	}
	return -1
}
