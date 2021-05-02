package bricks

func leastBricks(wall [][]int) int {
	ret := 0
	counter := map[int]int{}
	for _, row := range wall {
		sum := 0
		for _, brick := range row[:len(row)-1] {
			sum += brick
			_, ok := counter[sum]
			if !ok {
				counter[sum] = 0
			}
			counter[sum] += 1
			if counter[sum] > ret {
				ret = counter[sum]
			}
		}
	}
	return len(wall) - ret
}
