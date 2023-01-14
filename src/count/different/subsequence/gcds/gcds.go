package gcds

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func countDifferentSubsequenceGCDs(nums []int) int {
	max := 0
	for _, num := range nums {
		if max < num {
			max = num
		}
	}
	m := make([]bool, max+1)
	ret := 0
	for _, num := range nums {
		if !m[num] {
			m[num] = true
			ret++
		}
	}
	for i := 1; i <= max; i++ {
		// 找i的倍数，看是否多于1个，并且他们的排一起的最大公约数是i
		if !m[i] {
			g := 0
			for j := i * 2; j <= max; j += i {
				if m[j] {
					if g == 0 {
						g = j
					} else {
						g = gcd(g, j)
					}
				}
			}
			if g == i {
				ret++
			}
		}
	}
	return ret
}
