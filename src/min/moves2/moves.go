package moves2

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minMoves(nums []int, limit int) int {
	n := len(nums)
	if n == 2 {
		return 0
	}
	var m int = 2*1e5 + 1
	diff := make([]int, m)
	for i := 0; i*2+1 < n; i++ {
		a, b := nums[i], nums[n-1-i]
		s := a + b
		l := s - max(a, b) + 1
		r := s + max(limit-min(a, b), 0)
		diff[2] += 2
		diff[l]--
		diff[s]--
		if s+1 < m {
			diff[s+1]++
		}
		if r+1 < m {
			diff[r+1]++
		}
	}
	out := diff[2]
	temp := out
	for i := 3; i < m; i++ {
		temp += diff[i]
		if temp < out {
			out = temp
		}
	}
	return out
}
