package coordinate

import "math"

func bestCoordinate(towers [][]int, radius int) []int {
	x1 := 50
	x2 := 0
	y1 := 50
	y2 := 0
	for _, tower := range towers {
		if tower[0] < x1 {
			x1 = tower[0]
		}
		if tower[0] > x2 {
			x2 = tower[0]
		}
		if tower[1] < y1 {
			y1 = tower[1]
		}
		if tower[1] > y2 {
			y2 = tower[1]
		}
	}
	ret := make([]int, 2)
	max := 0
	radiusSquare := radius * radius
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			temp := 0
			for k := 0; k < len(towers); k++ {
				w := towers[k][0] - i
				h := towers[k][1] - j
				m := w*w + h*h
				if m <= radiusSquare {
					temp += int(float64(towers[k][2]) / (1.0 + math.Ceil(math.Sqrt(float64(m)))))
				}
			}
			if temp > max {
				ret[0] = i
				ret[1] = j
			}
		}
	}
	return ret
}
