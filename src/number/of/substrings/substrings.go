package substrings

import "math"

func numberOfSubstrings(s string) int {
	n := len(s)
	zeroIndices := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			zeroIndices = append(zeroIndices, i)
		}
	}
	var out int
	var j int
	for i := 0; i < n; i++ {
		m := n - i
		count := int(math.Sqrt(float64(m)))
		for j < len(zeroIndices) && zeroIndices[j] < i {
			j++
		}
		min := i
		var end int
		for k := 0; k <= count && end != n; k++ {
			if j+k < len(zeroIndices) {
				end = zeroIndices[j+k]
			} else {
				end = n
			}
			l := end - i
			if l >= k*k+k {
				var start int
				if i+k*k+k-1 < min {
					start = min
				} else {
					start = i + k*k + k - 1
				}
				out += end - start
			}
			min = end
		}
	}
	return out
}
