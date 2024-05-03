package average

func average(salary []int) float64 {
	min := salary[0]
	max := min
	n := len(salary)
	sum := min
	for i := 1; i < n; i++ {
		s := salary[i]
		if s < min {
			min = s
		} else if s > max {
			max = s
		}
		sum += s
	}
	return float64(sum-min-max) / float64(n-2)
}
