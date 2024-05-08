package plants

func wateringPlants(plants []int, capacity int) int {
	var ret int
	temp := capacity
	for i, plant := range plants {
		if temp >= plant {
			temp -= plant
			ret++
		} else {
			temp = capacity - plant
			ret += 2*i + 1
		}
	}
	return ret
}
