package arrays

const mod int = 1e9 + 7

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func idealArrays(n int, maxValue int) int {
	sieve := make([]int, maxValue+1)
	for i := 2; i <= maxValue; i++ {
		if sieve[i] == 0 {
			for j := i; j <= maxValue; j += i {
				sieve[j] = i
			}
		}
	}
	ps := make([][]int, maxValue+1)
	var maxP int
	for i := 2; i <= maxValue; i++ {
		x := i
		for x > 1 {
			// 质数筛
			p := sieve[x]
			var count int
			for x%p == 0 {
				count++
				x /= p
			}
			if count > maxP {
				maxP = count
			}
			ps[i] = append(ps[i], count)
		}
	}
	e1 := n + maxP
	c := make([][]int, e1)
	for i := 0; i < e1; i++ {
		c[i] = make([]int, maxP+1)
	}
	c[0][0] = 1
	for i := 1; i < e1; i++ {
		c[i][0] = 1
		e2 := min(i, maxP)
		for j := 1; j <= e2; j++ {
			// 组合数递推
			c[i][j] = (c[i-1][j-1] + c[i-1][j]) % mod
		}
	}
	var ret int64
	for i := 1; i <= maxValue; i++ {
		var temp int64 = 1
		for _, p := range ps[i] {
			// 可重复组合
			// 相同位置的数可相乘得到一个新的数，进而得到一个新的组合，由于两组数之间互质
			// 所以可以应用乘法原理
			// 每个序列不可全为1，所以是n-1个空位与p块夹板的空位之间，夹板之间的顺序无关的排列数
			temp = (temp * int64(c[n+p-1][p])) % int64(mod)
		}
		ret = (ret + temp) % int64(mod)
	}
	return int(ret)
}
