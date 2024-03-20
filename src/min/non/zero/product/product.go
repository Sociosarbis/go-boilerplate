package product

const mask int64 = 1e9 + 7

func minNonZeroProduct(p int) int {
	var num int64 = (1 << p) - 2
	return int((dfs(num%mask, num/2) * ((num + 1) % mask)) % mask)
}

func dfs(v int64, power int64) int64 {
	if power == 0 {
		return 1
	} else if power == 1 {
		return v
	}
	temp := dfs(v, power/2)
	temp = (temp * temp) % mask
	if power%2 == 1 {
		temp = (temp * v) % mask
	}
	return temp
}
