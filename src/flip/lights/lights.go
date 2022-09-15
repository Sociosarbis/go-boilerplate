package lights

func flipLights(n int, presses int) int {
	if n > 4 {
		n = 4
	}
	mask := 1<<n - 1
	m := map[int]bool{
		mask: true,
	}
	ops := []int{15, 5, 10, 9}
	for i := 0; i < presses; i++ {
		nextM := make(map[int]bool, 8)
		for v := range m {
			for _, op := range ops {
				nextM[(v^op)&mask] = true
			}
		}
		m = nextM
	}
	return len(m)
}
