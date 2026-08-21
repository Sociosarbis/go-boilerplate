package smallest2

func comb(n, k int) int {
	if k == 1 {
		return n
	}
	out := n
	for i := 2; i <= k; i++ {
		out = out * (n - i + 1) / 2
	}
	return out
}

func count(coin int, value int64) int64 {
	return value / int64(coin)
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func lcm(a int, b int) int {
	c := gcd(a, b)
	return a / c * b
}

func findKthSmallest(coins []int, k int) int64 {
	n := len(coins)
	pcs := make([]int, 0, n)
block:
	for i, coin := range coins {
		for j := 0; j < n; j++ {
			if j != i {
				if coin%coins[j] == 0 {
					continue block
				}
			}
		}
		pcs = append(pcs, coin)
	}
	n = len(pcs)
	// 构建子集的最小公倍数
	groups := make([][][2]int, n)
	groups[0] = make([][2]int, 0, n)
	for i, coin := range pcs {
		groups[0] = append(groups[0], [2]int{coin, i})
	}
	for i := 1; i < n; i++ {
		groups[i-1] = make([][2]int, 0, comb(n, i+1))
		for _, option := range groups[i-1] {
			value, s := option[0], option[1]
			for j := s + 1; j < n; j++ {
				groups[i] = append(groups[i], [2]int{lcm(value, pcs[j]), j})
			}
		}
	}
	k64 := int64(k)
	l := k64
	r := 25 * k64
	out := r
	for l <= r {
		mid := (l + r) / 2
		var temp int64
		// 容斥原理
		for i := 0; i < n; i += 2 {
			for _, option := range groups[i] {
				temp += count(option[0], mid)
			}
		}
		for i := 1; i < n; i += 2 {
			for _, option := range groups[i] {
				temp -= count(option[0], mid)
			}
		}
		if temp >= k64 {
			out = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return out
}
