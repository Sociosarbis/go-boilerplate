package numbers

const total int = 9
const mod int64 = 1e9 + 7

func dfs(n int64) int {
	if n < 14 {
		num := 5
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
	ret := int64(dfs(a))
	return int((((ret * ret) % mod) * int64(dfs(n-2*a))) % mod)
}

func countGoodNumbers(n int64) int {
	return dfs(n)
}
