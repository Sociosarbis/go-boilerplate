package equal

func canBeEqual(target []int, arr []int) bool {
	if len(target) != len(arr) {
		return false
	}
	m := make([]int, 1000)
	c := 0
	for _, num := range target {
		if m[num-1] == 0 {
			c++
		}
		m[num-1]++
	}
	for _, num := range arr {
		if m[num-1] == 0 {
			return false
		}
		m[num-1]--
		if m[num-1] == 0 {
			c--
		}
	}
	return c == 0
}
