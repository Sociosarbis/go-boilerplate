package fruits

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	var out int
loop:
	for _, fruit := range fruits {
		for j, basket := range baskets {
			if fruit <= basket {
				baskets[j] = 0
				continue loop
			}
		}
		out++
	}
	return out
}
