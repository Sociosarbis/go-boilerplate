package array

func isZeroArray(nums []int, queries [][]int) bool {
	n := len(nums)
	diffs := make([]int, n)
	for _, q := range queries {
		l, r := q[0], q[1]
		diffs[l]++
		if r+1 < n {
			diffs[r+1]--
		}
	}
	var temp int
	for i, num := range nums {
		temp += diffs[i]
		if num > temp {
			return false
		}
	}
	return true
}
