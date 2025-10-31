package numbers

func getSneakyNumbers(nums []int) []int {
	n := len(nums) - 2
	visited := make([]bool, n)
	out := make([]int, 0, 2)
	for _, num := range nums {
		if visited[num] {
			out = append(out, num)
		} else {
			visited[num] = true
		}
	}
	return out
}
