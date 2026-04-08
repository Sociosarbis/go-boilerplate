package queries

const mask int = 1e9 + 7

func xorAfterQueries(nums []int, queries [][]int) int {
	for _, query := range queries {
		l, r, k, v := query[0], query[1], query[2], query[3]
		for i := l; i <= r; i += k {
			nums[i] = (nums[i] * v) % mask
		}
	}
	var out int
	for _, num := range nums {
		out ^= num
	}
	return out
}
