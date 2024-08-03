package square

const offset byte = 97
const charCount int = 26
const maxSize int = 1e9 + 1

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func maxPointsInsideSquare(points [][]int, s string) int {
	m := make([]int, charCount)

	for i := 0; i < charCount; i++ {
		m[i] = maxSize
	}

	limit := maxSize

	for i, point := range points {
		x, y := point[0], point[1]
		var size int
		if abs(x) > abs(y) {
			size = abs(x)
		} else {
			size = abs(y)
		}
		index := s[i] - offset
		if size >= m[index] {
			if size-1 < limit {
				limit = size - 1
			}
		} else {
			if m[index] != maxSize {
				if m[index]-1 < limit {
					limit = m[index] - 1
				}
			}
			m[index] = size
		}
	}

	var ret int

	for _, size := range m {
		if size != maxSize && size <= limit {
			ret++
		}
	}
	return ret
}
