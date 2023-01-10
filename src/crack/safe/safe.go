package safe

import "fmt"

func dfs(s string, mask, n, k, target int, m map[int]struct{}) string {
	if len(m) == target {
		return s
	}
	mask = (mask << 4) & ((1 << (4 * n)) - 1)
	for i := 0; i < k; i++ {
		if _, ok := m[mask|i]; !ok {
			m[mask|i] = struct{}{}
			res := dfs(fmt.Sprintf("%s%d", s, i), mask|i, n, k, target, m)
			if len(res) != 0 {
				return res
			}
			delete(m, mask|i)
		}
	}

	return ""
}

func crackSafe(n int, k int) string {
	ret := ""
	mask := 0
	target := 1
	for i := 0; i < n; i++ {
		target *= k
		if i <= k {
			ret += fmt.Sprint(i)
			mask |= i << ((n - i - 1) * 4)
		} else {
			ret += fmt.Sprint(k - 1)
			mask |= (k - 1) << ((n - i - 1) * 4)
		}
	}
	m := map[int]struct{}{}
	m[mask] = struct{}{}
	return dfs(ret, mask, n, k, target, m)
}
