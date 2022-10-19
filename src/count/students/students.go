package students

func countStudents(students []int, sandwiches []int) int {
	groups := make([]int, 2)
	for _, student := range students {
		groups[student]++
	}
	for i, sandwich := range sandwiches {
		if groups[sandwich] == 0 {
			return len(sandwiches) - i
		}
		groups[sandwich]--
	}
	return 0
}
