package values

func gcdValues(nums []int, queries []int64) []int {
	var m int
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	cnt := make([]int64, m+1)
	// 每个数字的数量
	for _, num := range nums {
		cnt[num]++
	}
	// 可整除i的数字数量
	cnt[1] = int64(len(nums))
	for i := 2; i <= m; i++ {
		for j := i * 2; j <= m; j += i {
			cnt[i] += cnt[j]
		}
	}
	// 以i为公因子的数对对数
	for i := 1; i <= m; i++ {
		cnt[i] = cnt[i] * (cnt[i] - 1) / 2
	}
	// 以i为最大公因子的数对对数
	for i := m / 2; i > 0; i-- {
		for j := i * 2; j <= m; j += i {
			cnt[i] -= cnt[j]
		}
	}
	// 数量的前缀和
	for i := 2; i <= m; i++ {
		cnt[i] += cnt[i-1]
	}
	out := make([]int, len(queries))
	for i, query := range queries {
		l := 1
		r := m
		for l <= r {
			mid := (l + r) / 2
			if query+1 > cnt[mid] {
				l = mid + 1
			} else {
				if mid > 0 && query+1 > cnt[mid-1] {
					out[i] = mid
					break
				} else {
					r = mid - 1
				}
			}
		}
	}
	return out
}
