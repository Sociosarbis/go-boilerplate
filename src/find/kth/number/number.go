package number

func dfs(temp int, n int, k *int, count int, acc int) int {
	rest := 0
	if n/count >= temp {
		rest = n - temp*count + 1
		if rest > count {
			rest = count
		}
	}
	if acc+rest < *k {
		*k -= acc + rest
		return 0
	}
	*k -= 1
	if *k == 0 {
		return temp
	}

	if temp < 1e9 {
		temp *= 10
		count /= 10
		acc -= count
		for i := 0; i <= 9; i++ {
			next := temp + i
			if next > n {
				break
			}
			ret := dfs(next, n, k, count, acc)
			if ret != 0 {
				return ret
			}
		}
	}
	return 0
}

func findKthNumber(n int, k int) int {
	temp := n / 10
	base := 1
	acc := 0
	for temp != 0 {
		temp /= 10
		acc += base
		base *= 10
	}
	for i := 1; i <= 9; i++ {
		ret := dfs(i, n, &k, base, acc)
		if ret != 0 {
			return ret
		}
	}
	return 0
}
