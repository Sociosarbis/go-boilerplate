package burgers

func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	temp := tomatoSlices - 2*cheeseSlices
	if temp >= 0 && temp&1 == 0 {
		x := temp / 2
		y := cheeseSlices - x
		if y >= 0 {
			return []int{x, y}
		}
	}
	return nil
}
