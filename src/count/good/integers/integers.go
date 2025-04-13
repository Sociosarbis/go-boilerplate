package integers

type empty struct{}

func setMask(mask int64, i, count int) int64 {
	return mask | int64(count<<(4*i))
}

func factorial(n int) int {
	ret := 1
	for i := 2; i <= n; i++ {
		ret *= i
	}
	return ret
}

func dfs(nums []int, visited map[int64]empty, n, i, k int) int64 {
	if i < len(nums) {
		var ret int64
		for j := 0; j < 10; j++ {
			nums[i] = j
			ret += dfs(nums, visited, n, i+1, k)
		}
		return ret
	}
	base := 1
	l := len(nums)
	var num int
	for i := 0; i < l; i++ {
		num += nums[i] * base
		base *= 10
	}
	if l == n/2 {
		i = l - 1
	} else {
		i = l - 2
	}
	for ; i >= 0; i-- {
		num += nums[i] * base
		base *= 10
	}
	if num%k != 0 {
		return 0
	}
	counter := [10]int{}
	for i := l - 2; i >= 0; i-- {
		counter[nums[i]] += 2
	}
	if l*2 == n {
		counter[nums[l-1]] += 2
	} else {
		counter[nums[l-1]]++
	}
	var mask int64
	for i := 0; i < 10; i++ {
		if counter[i] != 0 {
			mask |= setMask(mask, i, counter[i])
		}
	}
	if _, ok := visited[mask]; ok {
		return 0
	}
	visited[mask] = empty{}
	total := factorial(n)
	for i := 0; i < 10; i++ {
		if counter[i] > 1 {
			total /= factorial(counter[i])
		}
	}
	if counter[0] != 0 {
		total2 := factorial(n - 1)
		counter[0]--
		for i := 0; i < 10; i++ {
			if counter[i] > 1 {
				total2 /= factorial(counter[i])
			}
		}
		return int64(total - total2)
	}
	return int64(total)
}

func countGoodIntegers(n int, k int) int64 {
	l := n / 2
	if n&1 == 1 {
		l++
	}
	nums := make([]int, l)
	visited := map[int64]empty{}
	var ret int64
	for i := 1; i < 10; i++ {
		nums[0] = i
		ret += dfs(nums, visited, n, 1, k)
	}
	return ret
}
