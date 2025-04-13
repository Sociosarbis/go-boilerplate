package numbers

const total int = 9
const mod int64 = 1e9 + 7

func dfs(n int64, i int) int {
	if n < 14 {
		var num int
		if i == 0 {
			num = 5
		} else {
			num = 4
		}
		var j int64
		ret := 1
		for ; j < n; j++ {
			ret *= num
			num = total - num
		}
		return ret
	}
	a := n / 2
	if a%2 == 1 {
		a--
	}
	ret := int64(dfs(a, i))
	return int((((ret * ret) % mod) * int64(dfs(n-2*a, i))) % mod)
}

func countGoodNumbers(n int64) int {
	return dfs(n, 0)
}
