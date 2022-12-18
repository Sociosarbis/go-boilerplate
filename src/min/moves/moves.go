package moves

func minMoves(nums []int, k int) int {
	prefixSums := make([]int, len(nums)+1)
	window := []int{}
	ret := 2147483647
	leftCount := k / 2
	var rightCount int
	if k&1 == 1 {
		rightCount = k / 2
	} else {
		rightCount = k/2 - 1
	}
	for i, num := range nums {
		if num == 1 {
			prefixSums[i+1] = prefixSums[i] + i
			window = append(window, i)
			if len(window) > k {
				window = window[1:]
			}
			if len(window) == k {
				mid := window[leftCount]
				left := (mid*2-1-leftCount)*leftCount/2 - (prefixSums[mid] - prefixSums[window[0]])
				var right int
				if rightCount != 0 {
					right = (prefixSums[i+1] - prefixSums[mid+1]) - (mid*2+1+rightCount)*rightCount/2
				}
				if left+right < ret {
					ret = left + right
				}
			}
		} else {
			prefixSums[i+1] = prefixSums[i]
		}
	}
	return ret
}
