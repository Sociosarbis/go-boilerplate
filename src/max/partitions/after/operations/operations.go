package operations

func findRightMost(prefixSums [][26]int, start, k int, a, b byte) int {
	l, r := start, len(prefixSums)-2
	out := start
	for l <= r {
		mid := (l + r) >> 1
		var count int
		var i byte
		for ; i < 26; i++ {
			temp := prefixSums[mid+1][i] - prefixSums[start][i]
			if i == a {
				temp--
			} else if i == b {
				temp++
			}
			if temp > 0 {
				count++
			}
		}
		if count <= k {
			if mid > out {
				out = mid
			}
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return out
}

func maxPartitionsAfterOperations(s string, k int) int {
	n := len(s)
	prefixSums := make([][26]int, n+1)
	for i, c := range s {
		var j rune
		for ; j < 26; j++ {
			if c == j+97 {
				prefixSums[i+1][j] = prefixSums[i][j] + 1
			} else {
				prefixSums[i+1][j] = prefixSums[i][j]
			}
		}
	}
	pref := make([]int, n+1)
	partition_start := make([]int, n)
	var mask, count, start int
	for i := 0; i < n; i++ {
		bit := 1 << (s[i] - 'a')
		if mask&bit == 0 {
			count++
			if count > k || count == 1 {
				pref[i+1] = pref[i] + 1
				mask = bit
				count = 1
				start = i
			} else {
				pref[i+1] = pref[i]
				mask |= bit
			}
		} else {
			pref[i+1] = pref[i]
		}
		partition_start[i] = start
	}
	mask, count = 0, 0
	suff := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		bit := 1 << (s[i] - 'a')
		if mask&bit == 0 {
			count++
			if count > k || count == 1 {
				suff[i] = suff[i+1] + 1
				mask = bit
				count = 1
			} else {
				suff[i] = suff[i+1]
				mask |= bit
			}
		} else {
			suff[i] = suff[i+1]
		}
	}
	out := pref[n]
	for i := 0; i < n; i++ {
		var j byte
		start = partition_start[i]
		for ; j < 26; j++ {
			if j+'a' != s[i] {
				r := findRightMost(prefixSums, start, k, s[i]-'a', j)
				var temp int
				// 假如i不能作为分割点
				if r >= i {
					temp = 1 + pref[start] + suff[r+1]
				} else {
					r2 := findRightMost(prefixSums, i, k, s[i]-'a', j)
					temp = 2 + pref[start] + suff[r2+1]
				}
				if temp > out {
					out = temp
				}
			}
		}
	}
	return out
}
