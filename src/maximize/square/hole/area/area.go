package area

import "sort"

func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {
	sort.Ints(hBars)
	sort.Ints(vBars)
	var h, v int
	if hBars[0] == 1 {
		hBars = hBars[1:]
	}
	if len(hBars) > 0 && hBars[len(hBars)-1]-2 == n {
		hBars = hBars[:len(hBars)-1]
	}
	if vBars[0] == 1 {
		vBars = vBars[1:]
	}
	if len(vBars) > 0 && vBars[len(vBars)-1]-2 == n {
		vBars = vBars[:len(vBars)-1]
	}
	temp := 1
	l := len(hBars)
	for i := 1; i < l; i++ {
		if hBars[i] == hBars[i-1]+1 {
			temp++
		} else {
			if temp > h {
				h = temp
			}
			temp = 1
		}
	}
	if temp > h {
		h = temp
	}
	l = len(vBars)
	temp = 1
	for i := 1; i < l; i++ {
		if vBars[i] == vBars[i-1]+1 {
			temp++
		} else {
			if temp > v {
				v = temp
			}
			temp = 1
		}
	}
	if temp > v {
		v = temp
	}
	if h > v {
		h = v
	}
	return (h + 1) * (h + 1)
}
