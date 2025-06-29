package seq

import "sort"

const mask int = 1e9 + 7

func qpow(n int, k int) int {
	out := 1
	for k != 0 {
		if k&1 == 1 {
			out = (out * n) % mask
			k--
		} else {
			n = (n * n) % mask
			k >>= 1
		}
	}
	return out
}

func sumGeo2Seq(n int) int {
	return (qpow(2, n) - 1)
}

func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	var l int
	n := len(nums)
	r := n - 1
	var out int
	for l <= r {
		for l <= r && nums[l]+nums[r] > target {
			r--
		}
		if l < r {
			out = (out + sumGeo2Seq(r-l)) % mask
		}
		if l <= r {
			out = (out + 1) % mask
		}
		l++
	}
	return out
}
