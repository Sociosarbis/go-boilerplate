package queries2

import (
	"math"
)

const mask int64 = 1e9 + 7

func ceil(n, k int) int {
	if n%k == 0 {
		return n / k
	}
	return n/k + 1
}

func qpow(n int64, k int) int64 {
	if n == 1 {
		return n
	}
	var out int64 = 1
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

func xorAfterQueries(nums []int, queries [][]int) int {
	buckets := map[[2]int][][2]int{}
	n := len(nums)
	threshold := int(math.Sqrt(float64(n)))
	for _, query := range queries {
		l, r, k, v := query[0], query[1], query[2], query[3]
		if k <= threshold {
			key := [2]int{k, l % k}
			var bucket [][2]int
			length := ceil(n, k)
			if b, ok := buckets[key]; ok {
				bucket = b
			} else {
				bucket = make([][2]int, length)
				for i := 0; i < length; i++ {
					bucket[i][0], bucket[i][1] = 1, 1
				}
			}
			s := (l - key[1]) / k
			e := (r - key[1]) / k
			bucket[s][0] = int((int64(bucket[s][0]) * int64(v)) % mask)
			if e+1 < length {
				bucket[e+1][1] = int((int64(bucket[e+1][1]) * int64(v)) % mask)
			}
			buckets[key] = bucket
		} else {
			for i := l; i <= r; i += k {
				nums[i] = int((int64(nums[i]) * int64(v)) % mask)
			}
		}
	}
	for key, bucket := range buckets {
		var mul int64 = 1
		for i, pair := range bucket {
			mul = (((mul * int64(pair[0])) % mask) * qpow(int64(pair[1]), int(mask)-2)) % mask
			index := i*key[0] + key[1]
			if index < n {
				nums[index] = int((int64(nums[index]) * mul) % mask)
			}
		}
	}
	var out int
	for _, num := range nums {
		out ^= num
	}
	return out
}
