package sums

import "sort"

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a < b {
		a, b = b, a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func gcdSum(nums []int) int64 {
	var mx int
	prefixGcd := make([]int, len(nums))
	for i, num := range nums {
		mx = max(num, mx)
		prefixGcd[i] = gcd(num, mx)
	}
	sort.Ints(prefixGcd)
	var l int
	var out int64
	r := len(prefixGcd) - 1
	for l < r {
		out += int64(gcd(prefixGcd[r], prefixGcd[l]))
		l++
		r--
	}
	return out
}
