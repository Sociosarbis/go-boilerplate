package peaks

func findPeaks(mountain []int) []int {
	n := len(mountain)
	j := 0
	for i := 1; i+1 < n; i++ {
		if mountain[i] > mountain[i-1] && mountain[i] > mountain[i+1] {
			mountain[j] = i
			j++
		}
	}
	return mountain[:j]
}
