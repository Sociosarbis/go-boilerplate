package distance2

func minimumDistance(nums []int) int {
	numToIndices := make([][]int, 101)
	for i, num := range nums {
		numToIndices[num] = append(numToIndices[num], i)
	}
	out := 100
	for i := 1; i <= 100; i++ {
		for j := len(numToIndices[i]) - 1; j >= 2; j-- {
			if numToIndices[i][j]-numToIndices[i][j-2] < out {
				out = numToIndices[i][j] - numToIndices[i][j-2]
			}
		}
	}
	if out == 100 {
		return -1
	}
	return out * 2
}
