package indices

func findKDistantIndices(nums []int, key int, k int) []int {
	var upper int = -1
	n := len(nums)
	out := make([]int, 0, n)
	for i, num := range nums {
		if num == key {
			l := i - k
			r := i + k
			if l <= upper {
				l = upper + 1
			}
			if r >= n {
				r = n - 1
			}
			for j := l; j <= r; j++ {
				out = append(out, j)
			}
			upper = r
		}
	}
	return out
}
