package drunk

func maxBottlesDrunk(numBottles int, numExchange int) int {
	out := numBottles
	for numBottles >= numExchange {
		out++
		numBottles += 1 - numExchange
		numExchange++
	}
	return out
}
