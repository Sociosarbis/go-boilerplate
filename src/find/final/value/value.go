package value

type empty struct{}

func findFinalValue(nums []int, original int) int {
	n := len(nums)
	m := make(map[int]empty, n)
	for _, num := range nums {
		if _, ok := m[num]; !ok {
			m[num] = empty{}
		}
	}
	for {
		if _, ok := m[original]; ok {
			original *= 2
		} else {
			break
		}
	}
	return original
}
