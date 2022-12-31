package seat

import "sort"

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	ret := 0
	for i := 0; i < len(seats); i++ {
		if seats[i] < students[i] {
			ret += students[i] - seats[i]
		} else if seats[i] > students[i] {
			ret += seats[i] - students[i]
		}
	}
	return ret
}
