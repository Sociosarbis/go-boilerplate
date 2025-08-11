package queries

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

func productQueries(n int, queries [][]int) []int {
	var i int
	nums := []int{}
	for n != 0 {
		if n&1 == 1 {
			nums = append(nums, i)
		}
		n >>= 1
		i++
	}
	prefix := make([]int, len(nums)+1)
	for i, num := range nums {
		prefix[i+1] = prefix[i] + num
	}
	out := make([]int, len(queries))
	for i, query := range queries {
		diff := prefix[query[1]+1] - prefix[query[0]]
		out[i] = qpow(2, diff)
	}
	return out
}
