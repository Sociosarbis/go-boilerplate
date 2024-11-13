package substrings2

func countKConstraintSubstrings(s string, k int, queries [][]int) []int64 {
	n := len(s)
	memo := make([]int, n)
	prefixSum := make([]int64, n+1)
	var j int
	var a int
	var b int
	for i := 0; i < n; i++ {
		for ; j < n; j++ {
			if s[j] == '0' {
				a++
			} else {
				b++
			}
			if a <= k || b <= k {
				continue
			} else {
				break
			}
		}
		if s[i] == '0' {
			a--
		} else {
			b--
		}
		memo[i] = j - 1
		prefixSum[i+1] = prefixSum[i] + int64(j-i)
		if j < n {
			if s[j] == '0' {
				a--
			} else {
				b--
			}
		}
	}
	ret := make([]int64, len(queries))
	for i, q := range queries {
		start, end := q[0], q[1]
		l := start
		r := end
		for l <= r {
			mid := (l + r) / 2
			if memo[mid] > end {
				r = mid - 1
			} else {
				if mid+1 <= r && memo[mid+1] <= end {
					l = mid + 1
				} else {
					l = mid
					break
				}
			}
		}
		if memo[l] > end {
			l--
		}
		ret[i] = prefixSum[l+1] - prefixSum[start] + int64(1+end-l)*int64(end-l)/2
	}
	return ret
}
