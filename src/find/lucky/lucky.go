package lucky

func findLucky(arr []int) int {
	n := len(arr)
	counter := make([]int, n+1)
	for _, num := range arr {
		if num <= n {
			counter[num]++
		}
	}
	for i := n; i >= 1; i-- {
		if counter[i] == i {
			return i
		}
	}
	return -1
}
