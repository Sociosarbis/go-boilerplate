package boxes

import "math"

func minimumBoxes(n int) int {
	i := 0
	temp := 0
	for temp+(2+i)*(i+1)/2 <= n {
		i++
		temp += (1 + i) * i / 2
	}
	if temp == n {
		return (1 + i) * i / 2
	} else {
		diff := n - temp
		// 第一层每加1个，立方块数都会加n
		r := int(math.Sqrt(float64(diff * 2)))
		for r*(r+1)/2 < diff {
			r++
		}
		return (1+i)*i/2 + r
	}
}
