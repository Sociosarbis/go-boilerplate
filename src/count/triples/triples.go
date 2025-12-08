package triples

import "math"

func countTriples(n int) int {
	var out int
	max := n * n
	for i := 1; i < n; i++ {
		a := i * i
		for j := 1; j < n; j++ {
			sum := a + j*j
			if sum > max {
				break
			}
			c := int(math.Sqrt(float64(sum)))
			if c*c == sum {
				out++
			}
		}
	}
	return out
}
