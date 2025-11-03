package cost5

func minCost(colors string, neededTime []int) int {
	var out, start int
	temp := neededTime[0]
	max := temp
	n := len(colors)
	for i := 1; i < n; i++ {
		if colors[i] != colors[i-1] {
			if i-start > 1 {
				out += temp - max
			}
			start, temp = i, neededTime[i]
			max = temp
		} else {
			if neededTime[i] > max {
				max = neededTime[i]
			}
			temp += neededTime[i]
		}
	}
	if n-start > 1 {
		out += temp - max
	}
	return out
}
