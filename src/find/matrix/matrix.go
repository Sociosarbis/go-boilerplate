package matrix

func findMatrix(nums []int) [][]int {
	var count int
	n := len(nums)
	counter := make([]int, n+1)
	for _, num := range nums {
		counter[num]++
		if counter[num] > count {
			count = counter[num]
		}
	}
	ret := make([][]int, count)
	for i := 0; i < count; i++ {
		for j := 1; j <= n; j++ {
			if counter[j] > 0 {
				counter[j]--
				ret[i] = append(ret[i], j)
			}
		}
	}
	return ret
}
