package array

func distinctDifferenceArray(nums []int) []int {
	n := len(nums)
	ret := make([]int, n)
	visited := make(map[int]struct{}, n)
	for i, num := range nums {
		if _, ok := visited[num]; !ok {
			visited[num] = struct{}{}
		}
		ret[i] = len(visited)
	}
	visited = make(map[int]struct{}, len(visited))
	for i := n - 1; i >= 0; i-- {
		ret[i] = ret[i] - len(visited)
		num := nums[i]
		if _, ok := visited[num]; !ok {
			visited[num] = struct{}{}
		}
	}

	return ret
}
