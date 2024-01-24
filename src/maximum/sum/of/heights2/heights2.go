package heights2

func maximumSumOfHeights(maxHeights []int) int64 {
	var ret int64

	var temp int64

	n := len(maxHeights)

	heights := make([]int, n)

	var right int64

	var left int64

	for i := 0; i < n; i++ {
		right -= int64(heights[i])
		if maxHeights[i] > heights[i] {
			heights[i] = maxHeights[i]
		}

		for j := i - 1; j >= 0; j-- {
			var temp int
			if heights[j+1] > maxHeights[j] {
				temp = maxHeights[j]
			} else {
				temp = heights[j+1]
			}
			if heights[j] != temp {
				left += int64(temp - heights[j])
				heights[j] = temp
			} else {
				break
			}
		}

		for j := i + 1; j < n; j++ {
			var temp int
			if heights[j-1] > maxHeights[j] {
				temp = maxHeights[j]
			} else {
				temp = heights[j-1]
			}
			if heights[j] != temp {
				right += int64(temp - heights[j])
				heights[j] = temp
			} else {
				break
			}
		}

		temp = left + right + int64(heights[i])
		if temp > ret {
			ret = temp
		}
		left += int64(heights[i])
	}

	return ret
}
