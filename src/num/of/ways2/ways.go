package ways2

const mod int64 = 1e9 + 7

func split(nums []int) ([]int, []int) {
	var left int
	root := nums[0]
	for _, num := range nums {
		if num < root {
			left++
		}
	}
	n := len(nums)
	a, b := make([]int, 0, left), make([]int, 0, n-left)
	for i := 1; i < n; i++ {
		if nums[i] < root {
			a = append(a, nums[i])
		} else {
			b = append(b, nums[i])
		}
	}
	return a, b
}

func safeProduct(a, b int) int {
	return int((int64(a) * int64(b)) % mod)
}

func qpow(n int64, k int) int64 {
	var out int64 = 1
	for k != 0 {
		if k&1 == 1 {
			out = (out * n) % mod
			k--
		} else {
			n = (n * n) % mod
			k >>= 1
		}
	}
	return out
}

func comb(m, n int) int {
	if n == 0 {
		return 1
	}
	numerator := m
	denominator := 1
	for i := 2; i <= n; i++ {
		numerator = safeProduct(numerator, m-i+1)
		denominator = safeProduct(denominator, i)
	}
	return safeProduct(numerator, int(qpow(int64(denominator), int(mod-2))))
}

func dfs(nums []int) int {
	if len(nums) < 3 {
		return 1
	}
	a, b := split(nums)
	c1, c2 := dfs(a), dfs(b)
	n, m := len(a), len(nums)-1
	return safeProduct(safeProduct(c1, c2), comb(m, n))
}

func numOfWays(nums []int) int {
	return dfs(nums) - 1
}
