package operations

func minNumberOperations(target []int) int {
	var out, top int
	for _, value := range target {
		if top != value {
			if top < value {
				out += value - top
			}
			top = value
		}
	}
	return out
}
