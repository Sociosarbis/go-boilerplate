package mountains

func stableMountains(height []int, threshold int) []int {
	n := len(height)
	var i int
	prev := height[0]
	for j := 1; j < n; j++ {
		if prev > threshold {
			height[i] = j
			i++
		}
		prev = height[j]
	}
	return height[:i]
}
