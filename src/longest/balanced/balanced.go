package balanced

func longestBalanced(nums []int) int {
	n := len(nums)
	var out int
	for i := 0; i < n; i++ {
		var temp int
		visited := make(map[int]bool, n)
		visited[nums[i]] = true
		if nums[i]&1 == 0 {
			temp++
		} else {
			temp--
		}
		for j := i + 1; j < n; j++ {
			if _, ok := visited[nums[j]]; !ok {
				visited[nums[j]] = true
				if nums[j]&1 == 0 {
					temp++
				} else {
					temp--
				}
			}
			if temp == 0 {
				if j-i+1 > out {
					out = j - i + 1
				}
			}
		}
	}
	return out
}
